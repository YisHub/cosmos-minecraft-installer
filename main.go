// Instalador del pack "Cosmos 26.3" para el launcher oficial de Minecraft en Windows.
// No incluye mods: los descarga de sus sitios oficiales y verifica el SHA-1 de cada uno.
package main

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	_ "embed"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//go:embed manifest.json
var manifestJSON []byte

type Manifest struct {
	PackName     string `json:"packName"`
	ProfileID    string `json:"profileId"`
	GameDirName  string `json:"gameDirName"`
	Minecraft    string `json:"minecraft"`
	FabricLoader string `json:"fabricLoader"`
	JavaArgs     string `json:"javaArgs"`
	Server       *struct {
		Name string `json:"name"`
		IP   string `json:"ip"`
	} `json:"server"`
	Files []PackFile `json:"files"`
}

type PackFile struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
	File string `json:"file"`
	URL  string `json:"url"`
	SHA1 string `json:"sha1"`
}

// Archivo donde se anotan los archivos que puso el instalador, para borrar
// versiones viejas al actualizar sin tocar mods que el usuario agregó a mano.
const installedListName = ".cosmos-installed.json"

var client = &http.Client{Timeout: 5 * time.Minute}

func main() {
	err := run()
	if err != nil {
		fmt.Println()
		fmt.Println("ERROR:", err)
	}
	fmt.Println()
	fmt.Print("Presiona Enter para salir...")
	bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
}

func run() error {
	var m Manifest
	if err := json.Unmarshal(manifestJSON, &m); err != nil {
		return fmt.Errorf("manifest invalido: %w", err)
	}

	fmt.Printf("=== Instalador %s (Minecraft %s + Fabric) ===\n\n", m.PackName, m.Minecraft)
	fmt.Println("Cierra el launcher de Minecraft antes de continuar.")
	fmt.Print("Presiona Enter para instalar...")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println()

	appdata := os.Getenv("APPDATA")
	if appdata == "" {
		return errors.New("no se encontro la variable APPDATA (¿es Windows?)")
	}
	mcDir := filepath.Join(appdata, ".minecraft")
	if _, err := os.Stat(mcDir); err != nil {
		return fmt.Errorf("no existe %s: abre el launcher oficial de Minecraft una vez y vuelve a ejecutar el instalador", mcDir)
	}
	gameDir := filepath.Join(mcDir, m.GameDirName)

	versionID, err := installFabric(mcDir, m.Minecraft, m.FabricLoader)
	if err != nil {
		return fmt.Errorf("instalando Fabric: %w", err)
	}

	if err := installFiles(gameDir, m.Files); err != nil {
		return err
	}

	if err := writeIfMissing(filepath.Join(gameDir, "config", "iris.properties"),
		[]byte("enableShaders=true\nshaderPack="+shaderPackName(m.Files)+"\n")); err != nil {
		return fmt.Errorf("configurando shaders: %w", err)
	}

	if m.Server != nil {
		added, err := addServer(filepath.Join(gameDir, "servers.dat"), m.Server.Name, m.Server.IP)
		if err != nil {
			return fmt.Errorf("agregando el servidor: %w", err)
		}
		if added {
			fmt.Printf("[ok] Servidor \"%s\" (%s) agregado a la lista\n", m.Server.Name, m.Server.IP)
		}
	}

	if err := updateProfiles(mcDir, m, versionID, gameDir); err != nil {
		return fmt.Errorf("creando el perfil del launcher: %w", err)
	}

	fmt.Println()
	fmt.Println("Listo! Abre el launcher de Minecraft, elige el perfil")
	fmt.Printf("\"%s\" y pulsa Jugar.\n", m.PackName)
	return nil
}

// installFabric descarga el perfil de Fabric desde meta.fabricmc.net, igual que el instalador oficial.
func installFabric(mcDir, mc, loader string) (string, error) {
	id := fmt.Sprintf("fabric-loader-%s-%s", loader, mc)
	url := fmt.Sprintf("https://meta.fabricmc.net/v2/versions/loader/%s/%s/profile/json", mc, loader)
	data, err := fetch(url)
	if err != nil {
		return "", err
	}
	var check struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(data, &check); err != nil || check.ID != id {
		return "", fmt.Errorf("respuesta inesperada de Fabric (%q)", check.ID)
	}
	dir := filepath.Join(mcDir, "versions", id)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	if err := os.WriteFile(filepath.Join(dir, id+".json"), data, 0o644); err != nil {
		return "", err
	}
	// El instalador oficial de Fabric deja un .jar vacio junto al json.
	jar := filepath.Join(dir, id+".jar")
	if _, err := os.Stat(jar); err != nil {
		if err := os.WriteFile(jar, nil, 0o644); err != nil {
			return "", err
		}
	}
	fmt.Printf("[ok] Fabric Loader %s para Minecraft %s\n", loader, mc)
	return id, nil
}

func installFiles(gameDir string, files []PackFile) error {
	listPath := filepath.Join(gameDir, installedListName)
	var previous []string
	if data, err := os.ReadFile(listPath); err == nil {
		json.Unmarshal(data, &previous)
	}

	current := map[string]bool{}
	for _, f := range files {
		rel := filepath.Join(f.Dir, f.File)
		current[rel] = true
		dest := filepath.Join(gameDir, rel)
		if sum, err := fileSHA1(dest); err == nil && sum == f.SHA1 {
			fmt.Printf("[ok] %s (ya estaba)\n", f.Name)
			continue
		}
		fmt.Printf("[..] Descargando %s...\n", f.Name)
		data, err := fetch(f.URL)
		if err != nil {
			return fmt.Errorf("descargando %s: %w", f.Name, err)
		}
		sum := sha1.Sum(data)
		if hex.EncodeToString(sum[:]) != f.SHA1 {
			return fmt.Errorf("%s: el archivo descargado no coincide con el esperado (hash)", f.Name)
		}
		if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(dest, data, 0o644); err != nil {
			return err
		}
		fmt.Printf("[ok] %s\n", f.Name)
	}

	for _, rel := range previous {
		if !current[rel] {
			if os.Remove(filepath.Join(gameDir, rel)) == nil {
				fmt.Printf("[ok] Eliminada version vieja: %s\n", rel)
			}
		}
	}

	var list []string
	for rel := range current {
		list = append(list, rel)
	}
	data, _ := json.MarshalIndent(list, "", "  ")
	return os.WriteFile(listPath, data, 0o644)
}

func shaderPackName(files []PackFile) string {
	for _, f := range files {
		if f.Dir == "shaderpacks" {
			return f.File
		}
	}
	return ""
}

// updateProfiles agrega o actualiza el perfil en los archivos de perfiles del launcher.
// Conserva todo lo demas que haya en el archivo.
func updateProfiles(mcDir string, m Manifest, versionID, gameDir string) error {
	names := []string{"launcher_profiles.json", "launcher_profiles_microsoft_store.json"}
	updated := 0
	for i, name := range names {
		path := filepath.Join(mcDir, name)
		root := map[string]any{}
		data, err := os.ReadFile(path)
		switch {
		case err == nil:
			if err := json.Unmarshal(data, &root); err != nil {
				return fmt.Errorf("%s no es JSON valido: %w", name, err)
			}
		case errors.Is(err, os.ErrNotExist) && i == 0:
			// Se crea launcher_profiles.json si no existe ninguno.
		default:
			continue
		}
		profiles, _ := root["profiles"].(map[string]any)
		if profiles == nil {
			profiles = map[string]any{}
		}
		now := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		p, _ := profiles[m.ProfileID].(map[string]any)
		if p == nil {
			p = map[string]any{"created": now, "icon": "Grass"}
		}
		p["name"] = m.PackName
		p["type"] = "custom"
		p["lastVersionId"] = versionID
		p["gameDir"] = gameDir
		p["javaArgs"] = m.JavaArgs
		p["lastUsed"] = now
		profiles[m.ProfileID] = p
		root["profiles"] = profiles

		out, err := json.MarshalIndent(root, "", "  ")
		if err != nil {
			return err
		}
		if err := os.WriteFile(path, out, 0o644); err != nil {
			return err
		}
		updated++
	}
	fmt.Printf("[ok] Perfil \"%s\" creado en el launcher\n", m.PackName)
	if updated == 0 {
		return errors.New("no se pudo escribir ningun archivo de perfiles")
	}
	return nil
}

// addServer crea servers.dat con el servidor si no existe. Si ya existe no lo toca,
// para no perder los servidores que el usuario haya agregado.
func addServer(path, name, ip string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		fmt.Println("[ok] Lista de servidores ya existente (no se modifica)")
		return false, nil
	}
	var b bytes.Buffer
	str := func(s string) {
		binary.Write(&b, binary.BigEndian, uint16(len(s)))
		b.WriteString(s)
	}
	b.WriteByte(10) // TAG_Compound raiz
	str("")
	b.WriteByte(9) // TAG_List "servers" de compounds
	str("servers")
	b.WriteByte(10)
	binary.Write(&b, binary.BigEndian, int32(1))
	b.WriteByte(8) // TAG_String name
	str("name")
	str(name)
	b.WriteByte(8) // TAG_String ip
	str("ip")
	str(ip)
	b.WriteByte(0) // fin del servidor
	b.WriteByte(0) // fin de la raiz
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return false, err
	}
	return true, os.WriteFile(path, b.Bytes(), 0o644)
}

func writeIfMissing(path string, data []byte) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func fetch(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "cosmos-minecraft-installer (github.com/YisHub/cosmos-minecraft-installer)")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: HTTP %d", strings.SplitN(url, "?", 2)[0], resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func fileSHA1(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

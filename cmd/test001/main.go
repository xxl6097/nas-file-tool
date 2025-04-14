package main

import (
	"encoding/json"
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/xxl6097/go-frp-panel/cmd"
	"github.com/xxl6097/go-frp-panel/internal/frps"
	"os"
	"path/filepath"
)

func main() {
	cmd.Execute(func() error {
		temp := os.TempDir()
		temp = filepath.Join(temp, "frps", "logs")
		err := os.MkdirAll(temp, 0755)
		if err != nil {
			fmt.Println(err)
		}
		cfg := &v1.ServerConfig{
			BindPort: 6000,
			BindAddr: "0.0.0.0",
			WebServer: v1.WebServerConfig{
				Addr:     "0.0.0.0",
				Port:     7200,
				User:     "admin",
				Password: "admin",
			}, HTTPPlugins: []v1.HTTPPluginOptions{
				{
					Name: "frps-panel",
					Addr: fmt.Sprintf("%s:%d", "0.0.0.0", 7200),
					Path: "/handler",
					Ops:  []string{"Login", "NewWorkConn", "NewUserConn", "NewProxy", "Ping"},
				},
			},
			Log: v1.LogConfig{
				To:      filepath.Join(temp, "frps.log"),
				MaxDays: 15,
			},
		}
		frps.Test(&frps.CfgModel{
			Frps: *cfg,
		})
		content, _ := json.Marshal(cfg)
		svv, err := frps.NewFrps(content, nil)
		if err != nil {
			return err
		}
		svv.Run()
		return nil
	})
}
if err := os.WriteFile(cfgPath, utils.ObjectToTomlText(cfg), 0o600); err != nil {
			glog.Warnf("write content to frpc config file error: %v", err)
		}
	}

	fmt.Println(cfgPath)
	fmt.Println(string(utils.ObjectToTomlText(cfg)))
	cls, err := frpc.NewFrpc(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("http://localhost:%d\n", cfg.WebServer.Port)
	cls.Run()
}
erverError)
		return
	}
	defer chunkFile.Close()

	_, err = io.Copy(chunkFile, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing chunk file: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Chunk uploaded successfully"})
}

func handleMergeChunks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fileId := r.FormValue("fileId")
	totalChunksStr := r.FormValue("totalChunks")
	totalChunks, err := strconv.Atoi(totalChunksStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid total chunks: %v", err), http.StatusBadRequest)
		return
	}

	originalFileName := r.FormValue("originalFileName")
	mergedFilePath := filepath.Join(uploadDir, originalFileName)
	mergedFile, err := os.Create(mergedFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating merged file: %v", err), http.StatusInternalServerError)
		return
	}
	defer mergedFile.Close()

	for i := 0; i < totalChunks; i++ {
		chunkFileName := fmt.Sprintf("%s_%d", fileId, i)
		chunkFilePath := filepath.Join(uploadDir, chunkFileName)
		chunkFile, err := os.Open(chunkFilePath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error opening chunk file: %v", err), http.StatusInternalServerError)
			return
		}
		_, err = io.Copy(mergedFile, chunkFile)
		chunkFile.Close()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error merging chunk file: %v", err), http.StatusInternalServerError)
			return
		}
		os.Remove(chunkFilePath)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "File merged successfully"})
}

func serveIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

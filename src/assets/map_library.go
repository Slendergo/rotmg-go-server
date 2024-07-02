package assets

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var GlobalMapLibrary *MapLibrary

type CachedMapData struct {
	Width  int32
	Height int32
	Tiles  []CachedMapTile
}

type CachedMapTile struct {
	Ground int32
	Object int32
	Region int32
}

type JsonMapData struct {
	Data   []byte     `json:"data"`
	Dict   []JsonDict `json:"dict"`
	Width  int        `json:"width"`
	Height int        `json:"height"`
}

type JsonDict struct {
	Ground  string       `json:"ground"`
	Objs    []JsonObject `json:"objs"`
	Regions []JsonObject `json:"regions"`
}

type JsonObject struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type MapLibrary struct {
	maps map[string]*CachedMapData
}

func NewMapLibrary() *MapLibrary {
	return &MapLibrary{
		maps: make(map[string]*CachedMapData),
	}
}

func (ml *MapLibrary) GetMapData(name string) *CachedMapData {
	mapData, ok := ml.maps[name]
	if ok {
		return mapData
	}
	return nil
}

func (ml *MapLibrary) ProcessFiles(directory string) error {
	startTime := time.Now()

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if filepath.Ext(path) == ".jm" {
				err := ml.processJMFile(directory, path)
				if err != nil {
					return fmt.Errorf("error processing XML file %s: %v", path, err)
				}
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	duration := time.Since(startTime)

	fmt.Printf("Maps's cached in %s\n", duration)
	return nil
}

func (ml *MapLibrary) processJMFile(basePath string, filePath string) error {

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var jsonMapData JsonMapData
	err = json.Unmarshal(data, &jsonMapData)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	relPath, err := filepath.Rel(basePath, filePath)
	if err != nil {
		return fmt.Errorf("error getting relative path: %v", err)
	}

	typeString := strings.ReplaceAll(relPath, "\\", "/")

	decompressedData, err := ml.decompressData(jsonMapData.Data)
	if err != nil {
		return err
	}

	rdr := bytes.NewReader(decompressedData)

	cachedMapData := &CachedMapData{
		Width:  int32(jsonMapData.Width),
		Height: int32(jsonMapData.Height),
		Tiles:  make([]CachedMapTile, jsonMapData.Width*jsonMapData.Height),
	}

	for i := 0; i < jsonMapData.Width*jsonMapData.Height; i++ {
		var tileIndex int16
		err := binary.Read(rdr, binary.BigEndian, &tileIndex)
		if err != nil {
			return err
		}

		dict := jsonMapData.Dict[tileIndex]

		tileType := GlobalXMLLibrary.GroundTypeFromId(dict.Ground)

		var objectType int32
		if len(dict.Objs) > 0 {
			objectType = GlobalXMLLibrary.TypeFromId(dict.Objs[0].Id)
		} else {
			objectType = -1
		}

		// todo enum from stirng?

		// var regionType int32
		// if len(dict.Regions) > 0 {
		// 	regionType = GlobalXMLLibrary.TypeFromId(dict.Objs[0].Id)
		// } else {
		// 	regionType = -1
		// }

		regionType := int32(-1)

		cachedMapData.Tiles[i] = CachedMapTile{
			Ground: tileType,
			Object: objectType,
			Region: regionType,
		}
	}

	ml.maps[typeString] = cachedMapData

	return nil
}

func (ml *MapLibrary) decompressData(data []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

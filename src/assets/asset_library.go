package assets

import (
	"errors"
	"fmt"
	assets "main/assets/xml"
	"main/utils"
	"os"
	"path/filepath"
	"time"

	"github.com/beevik/etree"
)

var GlobalAssetLibrary *AssetLibrary

type AssetLibrary struct {
	objectIdToType   map[string]int32
	objectTypeToId   map[int32]string
	objectProperties map[int32]*assets.XMLObjectProperties
	enemies          map[int32]*assets.XMLEnemy
	// players map[int32]*assets.XMLPlayer
	// items map[int32]*assets.XMLItem

	groundIdToType   map[string]int32
	groundTypeToId   map[int32]string
	groundProperties map[int32]*assets.XMLGround
}

func NewAssetLibrary() *AssetLibrary {
	return &AssetLibrary{
		objectIdToType:   make(map[string]int32),
		objectTypeToId:   make(map[int32]string),
		objectProperties: make(map[int32]*assets.XMLObjectProperties),
		enemies:          make(map[int32]*assets.XMLEnemy),

		groundIdToType:   make(map[string]int32),
		groundTypeToId:   make(map[int32]string),
		groundProperties: make(map[int32]*assets.XMLGround),
	}
}

func (al *AssetLibrary) ProcessFiles(directory string) error {
	startTime := time.Now()

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if filepath.Ext(path) == ".xml" {
				err := al.processXMLFile(path)
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

	fmt.Printf("XML's loaded in %s\n", duration)

	return nil
}

func (al *AssetLibrary) processXMLFile(filePath string) error {

	doc := etree.NewDocument()
	err := doc.ReadFromFile(filePath)
	if err != nil {
		return err
	}

	root := doc.SelectElement("Objects")
	if root != nil {

		for _, elem := range root.SelectElements("Object") {

			idName := elem.SelectAttrValue("id", "Unknown")
			if idName == "Unknown" {
				return errors.New("unable to find id attribute on object")
			}

			intType := utils.ParseHexInt32(elem.SelectAttrValue("type", "-1"))
			if intType == -1 {
				return errors.New("unable to find type attribute on object")
			}

			xmlObjectProperties := assets.NewXMLObjectProperties()
			xmlObjectProperties.IdName = idName
			xmlObjectProperties.Type = int32(intType)
			xmlObjectProperties.Parse(elem)

			al.objectIdToType[xmlObjectProperties.IdName] = xmlObjectProperties.Type
			al.objectTypeToId[xmlObjectProperties.Type] = xmlObjectProperties.IdName
			al.objectProperties[xmlObjectProperties.Type] = xmlObjectProperties

			if xmlObjectProperties.IsPlayer {
				// todo
			}

			if xmlObjectProperties.IsItem {
				// todo
			}

			if xmlObjectProperties.IsEnemy {
				xmlEnemy := assets.NewXMLEnemy()
				xmlEnemy.IdName = idName
				xmlEnemy.Type = int32(intType)
				xmlEnemy.Parse(elem)
				al.enemies[xmlEnemy.Type] = xmlEnemy
			}
		}
		return nil
	}

	root = doc.SelectElement("Grounds")
	if root != nil {

		for _, elem := range root.SelectElements("Ground") {

			idName := elem.SelectAttrValue("id", "Unknown")
			if idName == "Unknown" {
				return errors.New("unable to find id attribute on ground")
			}

			intType := utils.ParseHexInt32(elem.SelectAttrValue("type", "-1"))
			if intType == -1 {
				return errors.New("unable to find type attribute on object")
			}

			xmlGround := assets.NewXMLGround()
			xmlGround.IdName = idName
			xmlGround.Type = int32(intType)
			xmlGround.Parse(elem)

			al.groundIdToType[xmlGround.IdName] = xmlGround.Type
			al.groundTypeToId[xmlGround.Type] = xmlGround.IdName
			al.groundProperties[xmlGround.Type] = xmlGround
		}
	}

	return nil
}

func (al *AssetLibrary) IdFromType(typ int32, defaultValue string) string {
	idName, ok := al.objectTypeToId[typ]
	if ok {
		return idName
	}
	return defaultValue
}

func (al *AssetLibrary) TypeFromId(idName string, defaultValue int32) int32 {
	typ, ok := al.objectIdToType[idName]
	if ok {
		return typ
	}
	return defaultValue
}

func (al *AssetLibrary) GetXMLObjectProperties(typ int32) *assets.XMLObjectProperties {
	xmlObjectProperties, ok := al.objectProperties[typ]
	if ok {
		return xmlObjectProperties
	}
	return nil
}

func (al *AssetLibrary) GroundIdFromType(typ int32, defaultValue string) string {
	idName, ok := al.groundTypeToId[typ]
	if ok {
		return idName
	}
	return defaultValue
}

func (al *AssetLibrary) GroundTypeFromId(idName string, defaultValue int32) int32 {
	typ, ok := al.groundIdToType[idName]
	if ok {
		return typ
	}
	return defaultValue
}

func (al *AssetLibrary) GetXMLGround(typ int32) *assets.XMLGround {
	xmlGround, ok := al.groundProperties[typ]
	if ok {
		return xmlGround
	}
	return nil
}

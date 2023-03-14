package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/patrickmn/go-cache"
)

func init() {
	gob.Register([]interface{}{})
	gob.Register(map[string]interface{}{})
}

func getSalahsTimes(uuid string, key string) (map[string]interface{}, error) {
	var decodedResponse map[string]interface{}

	// Open the cache file for reading and writing
	file, err := os.OpenFile("cache.gob", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize an empty map to store the cache items
	items := make(map[string]cache.Item)

	// Attempt to decode the cache items from the file
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&items)
	if err != nil && err != io.EOF {
		return nil, err
	}

	// Create a new cache instance from the items loaded from the file
	c := cache.NewFrom(cache.DefaultExpiration, cache.DefaultExpiration, items)

	// Check if the key is present in the cache
	if cachedResponse, found := c.Get(key); found {
		decodedResponse = cachedResponse.(map[string]interface{})
	} else {
		// If the key is not present in the cache, fetch it from the API
		response, err := http.Get("https://time.my-masjid.com/api/TimingsInfoScreen/GetMasjidTimings?GuidId=" + uuid)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		decodedResponse = make(map[string]interface{})
		err = json.NewDecoder(response.Body).Decode(&decodedResponse)
		if err != nil {
			return nil, err
		}

		// Create a new cache item for the key and value
		c.Set(key, decodedResponse, cache.DefaultExpiration)

		// Encode the updated cache items to the file
		err = file.Truncate(0)
		if err != nil {
			return nil, err
		}
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return nil, err
		}
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(c.Items())
		if err != nil {
			return nil, err
		}
	}

	return decodedResponse, nil
}

func playAzan() {
	azanDir := filepath.Join(".", "media", "azan")
	duaDir := filepath.Join(".", "media", "duaa")

	azanFiles, err := ioutil.ReadDir(azanDir)

	if err != nil {
		log.Fatal(err)
	}

	duaFiles, err := ioutil.ReadDir(duaDir)

	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	indexAzan := rand.Intn(len(azanFiles))
	indexDuaa := rand.Intn(len(duaFiles))

	cmd := exec.Command("play", filepath.Join(azanDir, azanFiles[indexAzan].Name()), filepath.Join(duaDir, duaFiles[indexDuaa].Name()))

	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	response, err := getSalahsTimes("8695b116-07e4-44a7-a8c6-df8f6006bdf7", "salah-times")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(response)
	}
}
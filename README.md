# PixivGoGo - Golang client for Pixiv

*It's still under early development. 
Only a small subset of APIs have been implemented.*

## Usage

Example:
```
import (
	"encoding/json"
	"github.com/sleepingpig/pixivgogo/pkg/pixivgogo"
)

func main() {
	client := pixivgogo.NewClient()
	err := client.Login("username", "password")
	if err != nil {
		log.Fatalf("failed to login: %v", err)
	}
    recommendFilter := &pixivgogo.RecommendIllustsFilter{
        Filter:                "for_android",
    }
    recommendedIllusts, err := client.IllustRecommend(recommendFilter)
    if err != nil {
        log.Fatalf("failed to get recommended illustration: %v", err)
    }
    recommendedIllustsJSON := jsonMarshal(recommendedIllusts)
    log.Printf("Recommended Illustrations: %s\n", recommendedIllustsJSON)
}

func jsonMarshal(src interface{}) string {
	jsonBytes, err := json.Marshal(src)
	if err != nil {
		log.Fatalf("failed to marshal: %v", err)
	}
	return string(jsonBytes)
}
```

## Tool
A command line tool is provided.  
You can build the tool as below:  
```
go build ./cmd/pixivtool/
```
To cross compile, add the environment variables, e.g.  
```
GOOS=windows GOARCH=386 go build ./cmd/pixivtool/
```

package lesson01

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

// LoadEngineConf - ゲーム設定ファイルを読み込みます。
func LoadEngineConf(path string) (*EngineConf, error) {

	// ファイル読込
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// debugPrintToml(fileData)

	// Toml解析
	binary := []byte(string(fileData))
	config := &EngineConf{}
	toml.Unmarshal(binary, config)

	return config, nil
}

// 外部プロセスとして呼び出される場合、まだロガーの設定が終わってないかも知れません。
// func debugPrintToml(fileData []byte) {
// 	// Out.Log.Trace("...Engine content=%s", string(fileData))

// 	// Toml解析
// 	tomlTree, err := toml.Load(string(fileData))
// 	if err != nil {
// 		panic(err)
// 	}
// 	Out.Log.Trace("...Engine Input:\n")
// 	Out.Log.Trace("...Engine Engine.Komi=%f\n", tomlTree.Get("Engine.Komi").(float64))
// 	Out.Log.Trace("...Engine Engine.BoardSize=%d\n", tomlTree.Get("Engine.BoardSize").(int64))
// 	Out.Log.Trace("...Engine Engine.MaxMoves=%d\n", tomlTree.Get("Engine.MaxMoves").(int64))
// 	Out.Log.Trace("...Engine Engine.BoardData=%s\n", tomlTree.Get("Engine.BoardData").(string))
// }
// func debugPrintConfig(config e.EngineConf) {
// 	Out.Log.Trace("...Engine Memory:\n")
// 	Out.Log.Trace("...Engine Profile.Name=%s\n", config.Profile.Name)
// 	Out.Log.Trace("...Engine Profile.Pass=%s\n", config.Profile.Pass)
// 	Out.Log.Trace("...Engine Engine.Komi=%f\n", config.Engine.Komi)
// 	Out.Log.Trace("...Engine Engine.BoardSize=%d\n", config.Engine.BoardSize)
// 	Out.Log.Trace("...Engine Engine.MaxMoves=%d\n", config.Engine.MaxMoves)
// 	Out.Log.Trace("...Engine Engine.MaxMoves=%s\n", config.Engine.BoardData)
// 	Out.Log.Trace("...Engine Engine.SentinelBoardMax()=%d\n", config.SentinelBoardMax())
// }

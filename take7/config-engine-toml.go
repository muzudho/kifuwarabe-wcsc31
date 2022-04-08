package take7

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
// 	// u.App.LogNotEcho.Trace("...Engine content=%s", string(fileData))

// 	// Toml解析
// 	tomlTree, err := toml.Load(string(fileData))
// 	if err != nil {
// 		panic(err)
// 	}
// 	u.App.LogNotEcho.Trace("...Engine Input:\n")
// 	u.App.LogNotEcho.Trace("...Engine Engine.Komi=%f\n", tomlTree.Get("Engine.Komi").(float64))
// 	u.App.LogNotEcho.Trace("...Engine Engine.BoardSize=%d\n", tomlTree.Get("Engine.BoardSize").(int64))
// 	u.App.LogNotEcho.Trace("...Engine Engine.MaxMoves=%d\n", tomlTree.Get("Engine.MaxMoves").(int64))
// 	u.App.LogNotEcho.Trace("...Engine Engine.BoardData=%s\n", tomlTree.Get("Engine.BoardData").(string))
// }
// func debugPrintConfig(config e.EngineConf) {
// 	u.App.LogNotEcho.Trace("...Engine Memory:\n")
// 	u.App.LogNotEcho.Trace("...Engine Profile.Name=%s\n", config.Profile.Name)
// 	u.App.LogNotEcho.Trace("...Engine Profile.Pass=%s\n", config.Profile.Pass)
// 	u.App.LogNotEcho.Trace("...Engine Engine.Komi=%f\n", config.Engine.Komi)
// 	u.App.LogNotEcho.Trace("...Engine Engine.BoardSize=%d\n", config.Engine.BoardSize)
// 	u.App.LogNotEcho.Trace("...Engine Engine.MaxMoves=%d\n", config.Engine.MaxMoves)
// 	u.App.LogNotEcho.Trace("...Engine Engine.MaxMoves=%s\n", config.Engine.BoardData)
// 	u.App.LogNotEcho.Trace("...Engine Engine.SentinelBoardMax()=%d\n", config.SentinelBoardMax())
// }

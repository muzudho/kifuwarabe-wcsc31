package take5

// EngineConf - Tomlファイル。
type EngineConf struct {
	Profile Profile
}

// Profile - [Profile] 区画。
type Profile struct {
	Name   string
	Author string
}

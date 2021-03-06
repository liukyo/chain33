// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chaincfg

var configMap = make(map[string]string)

func Register(name, cfg string) {
	if _, ok := configMap[name]; ok {
		panic("chain default config name " + name + " is exist")
	}
	configMap[name] = cfg
}

func Load(name string) string {
	return configMap[name]
}

func LoadAll() map[string]string {
	return configMap
}

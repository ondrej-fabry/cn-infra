# Tutorial: Adding Dependencies to Your Plugin

Your control plane agent will typically consist of one or more plugins that
contain the application logic and a bunch of Ligato plugins that will 
provide services to your application plugins, such as KV- Data Store adapters, 
message bus adapters, loggers or health monitors. This tutorial shows how to
add plugin dependencies to your plugins. 

The Ligato infrastructure uses the **dependency injection** design pattern to
manage plugin dependencies. In other words, dependencies on other plugins are
injected into your plugin when it's initialized. You should dependency injection
to manage other non-plugin dependencies as well - it is critical for creation
of mocks in unit tests. Without dependency injection it is almost impossible to
achieve good unit test coverage.

The 'github.com/ligato/cn-infra/infra` package contains the `PluginDeps` struct
for easy embedding into plugins which provides essentials for plugins: plugin name,
logging and config. 

It is defined as:

```go
type PluginDeps struct {
	PluginName
	Log logging.PluginLogger
	Cfg config.PluginConfig
}
```

And used by simply embedding into any plugin:

```go
type HelloWorld struct {
	infra.PluginDeps
}
```

The `PluginDeps` has `PluginName` embedded into it to automatically provide 
the `String()` method for naming plugin, thus avoiding the need to define it for plugin
that embeds the `PluginDeps`. The name is set using `SetName(name string)` method which comes from
the `PluginName`.

```go
p.SetName("helloworld")
```

The `PluginDeps` also has `Setup()` method which initializes `Log` and `Cfg` by 
using the name from `PluginName`.

These are typically called in constructor of a plugin.

```go
func NewHelloWorld() *HelloWorld {
	p := new(HelloWorld)
	p.SetName("helloworld")
	p.Setup()
	return p
}
```

Now when `Log` and `Cfg` are initialized we can use them.

The plugin logger, `Log` can be used to log some message at different log levels.

```go
func (p *HelloWorld) Init() error {
	p.Log.Info("System ready.")
	p.Log.Warn("Problems found!")
	p.Log.Error("Errors encountered!")
}
```

The plugin config, `Cfg` can be used to load configuration from file using YAML format.

By default the name of the config file will be derived from the plugin name with extension `.conf`.
In our case, it will be `helloworld.conf`.

```go
type Config struct {
	MyValue int `json:"my-value"`
}

func (p *HelloWorld) Init() error {
	cfg := new(Config)
	found, err := p.Cfg.LoadValue(cfg)
	// ...
}
```

If the config file is not found, the `LoadValue` will return false. 
On parsing issues it will return an error.

Complete working example can be found at [examples/tutorials/02_plugin-deps](https://github.com/ligato/cn-infra/blob/master/examples/tutorials/02_plugin-deps).
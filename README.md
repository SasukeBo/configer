# Configer

yaml config file loader

## Usage

### Custom setting
configer support three config file: application, development, production
files default name is: app.yaml, dev.yaml and prod.yaml.
you can set your custom config file name.

```golang
// custom your application config file name
configer.SetEntryConfigFileName("main.yaml")
// custom your development config file name
configer.SetDevelopmentConfigFileName("development.yaml")
// custom your production config file name
configer.SetProductionConfigFileName("production.yaml")
```

configer consider the default configuration files located at project `./conf` directory.
you can set your custom conf directory by `SetConfigFileDir` function:

```golang
configer.SetConfigFileDir("./")
// or
configer.SetConfigFileDir("./system")
```

> Note that if you made a custom configuration, a reload inneed!
```golang
err := configer.ReloadConfig()
if err != nil {
  log.Fatal(err)
}
```

### Get config

```golang
// get a string config value
stringValue, err := Config.GetString("string_key")
// get a int config value
intValue, err := Config.GetInt("int_key")
// get a bool config value
boolValue, err := Config.GetBool("bool_key")
// get a float config value
floatValue, err := Config.GetFloat("float_key")
```

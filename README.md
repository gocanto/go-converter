## About the library

![tests workflow](https://github.com/voyago/converter/actions/workflows/test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/voyago/converter)](https://goreportcard.com/report/voyago/converter)
[![Go Reference](https://pkg.go.dev/badge/github.com/voyago/converter.svg)](https://pkg.go.dev/github.com/voyago/converter)

The converter is a data-agnostic currency conversion library that allows you to perform conversions based on a given
currency and price value pair.

## Installation

This library is based on [GO](https://golang.org). So before using it, make sure you have it installed in your machine.

Once you have done this, you will be able to pull this library in by typing the following command in your terminal within
your go application or library.

First, you need to ope you terminal and cd into your project directory, like so:

```shell
cd $HOME/your-project-rote-folder
```

After doing so, you would be able to type the following command to get this library installed. For instance,

```shell
go get github.com/voyago/converter
```

> Note: Make sure you have properly configured you go project.

## How do the data layer adapters work?

The converter is a driver-based library that abstracts away its data supply implementation using different adapters.
Therefore, you will not need to depend on a specific resource to retrieve given currencies' exchange rate.

We currently ship with [only one data adapter](https://currencylayer.com/), but we intend to add more as we need it. So,
please, do feel free to open a PR/Issue with a choice of your preference for us to discuss/geek about it.

### How do I create my own Driver?

Creating your own driver is easy since you only need to fulfill the given [store handler](https://github.com/voyago/converter/blob/main/pkg/store/handler/handler.go) interface.

You will be able to see a full example by inspecting the [currency layer handler](https://github.com/voyago/converter/blob/main/pkg/store/handler/currencyLayer/handler.go#L21).
Also, if you are feeling curious, you can see how the store is able to [resolve drivers](https://github.com/voyago/converter/blob/main/pkg/store/store.go#L42) based on an incoming request.

Moreover, this library ships with two different Currency Layer [implementations](https://github.com/voyago/converter/tree/main/pkg/store/handler/currencyLayer)
since we had to mock handlers detail at the [unit testing](https://github.com/voyago/converter/blob/main/tests/unit/store/handler/currencyLayer_test.go#L19) level.

## Documentation

* [Environment](#Environment)
* [Models](#Models)
* [Converter](#Converter)
* [Supported currencies](#Supported-currencies)
* [See consumers example](https://github.com/voyago/converter-tests)
* [Test suit](https://github.com/voyago/converter/tree/main/tests)

### Environment

The environment is controlled by the [env package](https://github.com/voyago/converter/tree/main/environment) shipped within
this library. It will allow you to choose what driver you would like to use when working in any given environment (staging, local, production, etc.)

It will also require you to specify your consumer app root folder name in order for it to resolve your env file path. [See example](https://github.com/voyago/converter-tests/blob/main/main.go)

### Models

Models represent the abstraction between the converter and the real world.

The converter utilizes the currency and currencies collection information to perform the required operations to find conversion
rates between currencies.

- **Currency:** It holds the currencies [body](https://github.com/voyago/converter/blob/main/pkg/model/currency.go) information.
- **Currencies:** It is a currency collection used to store information about all our [supported currencies](https://github.com/voyago/converter/blob/main/pkg/store/blueprint/currencies.go). It also serves as the main converter entry point to resolve currencies rate.
- **Price:** It serves as a conversion blueprint storing all the required information used to perform a given conversion operation

### Converter

The converter is the main package shipped by this library. It is in charge of performing conversion operations for a given
price and  destination currency.

This functionality is given in two different ways, based on a given [abstract store](#Store-based-Converter) or as a [standalone function](#Standalone-based-Converter).

### Store based Converter

The store is a fetching mechanism that is able to find currencies exchange rates for any available driver through a unified
interface. [See example](https://github.com/voyago/converter/blob/main/tests/unit/conversion/converter_test.go#L60-L84)

### Standalone based Converter

This standalone function allows you to perform currencies conversions bypassing the store implementation. Thus, you are
given all the necessary freedom to build the conversion request. [See example](https://github.com/voyago/converter/blob/main/tests/unit/conversion/converter_test.go#L12-L58)

### Supported currencies

We work with a predefined currencies [blueprint](https://github.com/voyago/converter/blob/main/resources/currencies.json)
that allows us to map given exchanges rates on demand. These rates are the ones used by our converter to perform operations.

## Contributing

Please, feel free to fork this repository to contribute to it by submitting a functionalities/bugs-fixes pull request to enhance it.

## License

Please see [License File](https://github.com/voyago/converter/blob/main/LICENSE) for more information.

## How can I thank you?

There are many ways you would be able to support my open source work. There is not a right one to choose, so the choice is yours.

Nevertheless :grinning:, I would propose the following

- :arrow_up: Follow me on [Twitter](https://twitter.com/gocanto).
- :star: Start the repository.
- :handshake: Open a pull request to fix/improve the codebase.
- :writing_hand: Open a pull request to improve the documentation.
- :coffee: Buy me a [coffee](https://github.com/sponsors/gocanto)?

> Thank you for reading this far. :blush:

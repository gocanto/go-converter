## Converter

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

## How do I use it?

WIP; :bowtie:

If you are feeling curious, you can check the whole [test suit](https://github.com/voyago/converter/tree/main/tests/unit/conversion) to find out how this library works?

Moreover, we have published a [playground](https://github.com/voyago/converter-tests) to demonstrate how this library can be consumed.

## Documentation
* [Models](#Models)

## Models

Models abstract the information related to currencies, plus add some convenient way to handle the types used by
the converter.

## Contributing

Please, feel free to fork this repository to contribute to it by submitting a functionalities/bugs-fixes pull request to enhance it.

## License

Please see [License File](https://github.com/voyago/converter/blob/main/LICENSE) for more information.

## How can I thank you?

Why not star this GitHub repository and share its link on your social network?

> Don't forget to [follow me on twitter](https://twitter.com/gocanto)!

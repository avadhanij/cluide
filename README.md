# Cluide - Command Line AI Guide

Cluide is a simple terminal utility to help you easily send queries to popular AI platforms and get answers
without having to leave the terminal.

# Support

Cluide can currently work with

1. OpenAI ChatGPT

# Usage

## Config

The basic requirement to interact with any AI platform is an API key. The API keys can either be set as environment variables or
be provided as part of the cluide-config YAML file. The file must be placed in your home directory - `$HOME/.cluide-config.yml`

Example file

```
openai:
    api_key: <insert key here>
```

## ChatGPT

To use cluide with ChatGPT, you will first need to get an API key from OpenAI developer platform - https://platform.openai.com/

Once you have the key, it will need to be sourced through an environment variable - `OPENAI_API_KEY`

After that, cluide can simply be invoked with your query.

E.g.: `cluide ask-chat "Tell me a haiku"`

The default model used is `gpt-4o-mini`. It can be overridden using `--model` flag.

E.g.: `cluide ask-chat "Tell me a haiku" --model gpt-4o`

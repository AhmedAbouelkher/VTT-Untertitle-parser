# VTT Untertitle parser

I am learning German 😃. this is a handy tool to parse the untertitles of the movies and series that I watch and translate them to English.

> *Caution*: Use it at your own risk. It may not be accurate. It is just a tool to help you to learn a language. I'm not responsible for any mistakes or errors.

## How to use it?

- Clone the repository
- run `go mod tidy`
- run `go run main.go -src <SOURCE_LANGUAGE> -dst <DESTINATION_LANGUAGE> <LOCAL_FILE_PATH>`
- You will find a new file with the same name but starts with `translated_` in the same directory as the original file.

### Example

```shell
go run main.go -src "de" -dst "en" ./data/untertitles.vtt
```

## How to get the untertitles?

Use your imagination. I'm not responsible for any illegal actions.

## How to contribute?

- Fork the repository
- Create a new branch
- Make your changes
- Create a pull request

## Credits

- [schollz/progressbar](https://github.com/schollz/progressbar) - A really cool progress bar for your terminal.
- [martinlindhe/subtitles](https://github.com/martinlindhe/subtitles) - A really cool library to parse subtitles.
- [ssut/py-googletrans#issue-268](https://github.com/ssut/py-googletrans/issues/268)

## License

[MIT](./LICENSE)

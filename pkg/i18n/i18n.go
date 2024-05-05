package i18n

import (
	"embed"
	"encoding/json"
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type I18n struct {
	opts     Options
	bundle   *i18n.Bundle
	localize *i18n.Localizer
	lang     language.Tag
}

func New(options ...func(*Options)) (rp *I18n) {
	opts := getOptionsOrSetDefault(nil)
	for _, f := range options {
		f(opts)
	}
	bundle := i18n.NewBundle(opts.language)
	localize := i18n.NewLocalizer(bundle, opts.language.String())
	switch opts.format {
	case "toml":
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	case "json":
		bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	default:
		bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	}
	rp = &I18n{
		opts:     *opts,
		bundle:   bundle,
		localize: localize,
		lang:     opts.language,
	}
	for _, item := range opts.files {
		rp.Add(item)
	}
	rp.AddFS(opts.fs)
	return
}

// Select can change language.
func (i I18n) Select(lang language.Tag) *I18n {
	if lang.String() == "und" {
		lang = i.opts.language
	}
	return &I18n{
		opts:     i.opts,
		bundle:   i.bundle,
		localize: i.localize,
		lang:     lang,
	}
}

// Language returns the current language.
func (i I18n) Language() language.Tag {
	return i.lang
}

// LocalizeT returns the localized message for the given message.
func (i I18n) LocalizeT(message *i18n.Message) (rp string) {
	if message == nil {
		return ""
	}
	var err error
	rp, err = i.localize.Localize(&i18n.LocalizeConfig{
		DefaultMessage: message,
	})
	if err != nil {
		rp = message.ID
	}
	return
}

// LocalizeE returns the localized message for the given message.
func (i I18n) LocalizeE(message *i18n.Message) string {
	return i.LocalizeT(message)
}

// T localizes the message with the given ID and returns the localized string.
// It uses the LocalizeT method to perform the translation.
func (i I18n) T(id string) (rp string) {
	return i.LocalizeT(&i18n.Message{ID: id})
}

// E is a wrapper for T that converts the localized string to an error type and returns it.
func (i I18n) E(id string) error {
	return errors.New(i.T(id))
}

func (i I18n) Add(file string) {
	info, err := os.Stat(file)
	if err != nil {
		return
	}
	if info.IsDir() {
		filepath.Walk(file, func(path string, fi os.FileInfo, errBack error) (err error) {
			if !fi.IsDir() {
				i.bundle.MustLoadMessageFile(path)
			}
			return
		})
	} else {
		i.bundle.MustLoadMessageFile(file)
	}
}

func (i *I18n) AddFS(fs embed.FS) {
	files := readFS(fs, ".")
	for _, name := range files {
		i.bundle.LoadMessageFileFS(fs, name)
	}
}

func readFS(fs embed.FS, dir string) (rp []string) {
	rp = make([]string, 0)
	dirs, err := fs.ReadDir(dir)
	if err != nil {
		return
	}
	for _, item := range dirs {
		name := dir + string(os.PathSeparator) + item.Name()
		if dir == "." {
			name = item.Name()
		}
		if item.IsDir() {
			rp = append(rp, readFS(fs, name)...)
		} else {
			rp = append(rp, name)
		}
	}
	return
}

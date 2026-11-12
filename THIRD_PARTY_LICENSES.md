# Third-Party Licenses

MiyooPod itself is licensed under the [MIT License](LICENSE). It bundles and
depends on the following third-party components, each distributed under its own
license. Copyright and license notices for these components are retained by
their respective authors.

## Bundled Native Libraries (`libs/`)

| Component | License |
| --- | --- |
| SDL2 | zlib License |
| SDL2_mixer | zlib License |
| zlib (`libSDL2_z`) | zlib License |
| json-c (`libSDL2_json-c`) | MIT License |
| libmpg123 | LGPL 2.1 |
| dr_flac (statically linked in SDL2_mixer) | Public Domain (Unlicense) / MIT-0 |
| stb_vorbis (statically linked in SDL2_mixer) | Public Domain (Unlicense) / MIT |

### Note on libmpg123 (LGPL 2.1)

libmpg123 is licensed under the GNU Lesser General Public License, version 2.1.
It is dynamically linked (as a shared `.so`), which satisfies the LGPL
requirement that users be able to replace or relink the library. Its license
text and source are available from https://www.mpg123.de/.

## Bundled Fonts (`assets/`)

| Font | File | License |
| --- | --- |
| Inter Regular | `ui_font.ttf` | SIL Open Font License 1.1 |
| DejaVu Sans | `ui_font_ar.ttf` | Bitstream Vera License / public-domain DejaVu changes |

DejaVu Sans is used when rendering Arabic (and other RTL) track metadata so
glyphs and presentation forms are available. Inter remains the default Latin UI
font.

Arabic letter-joining tables in `src/arabic_alphabet.go` are adapted from
[goarabic](https://github.com/01walid/goarabic) (MIT License).

## Go Module Dependencies

| Module | License |
| --- | --- |
| github.com/dhowden/tag | BSD-2-Clause |
| github.com/fogleman/gg | MIT License |
| github.com/google/uuid | BSD-3-Clause |
| github.com/skip2/go-qrcode | MIT License |
| golang.org/x/image | BSD-3-Clause |
| github.com/golang/freetype | FreeType License / GPLv2 |

Full license texts for each dependency are available in their respective source
repositories.

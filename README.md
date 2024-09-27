# Roast (Regal's Optimized AST)

Roast is an optimized JSON format for [Rego](https://www.openpolicyagent.org/docs/latest/policy-language/) ASTs, as well
as some common utilities for working with it.

Roast is used by [Regal](https://docs.styra.com/regal), where the JSON representation of Rego's AST is used input for
static analysis [performed by Rego itself](https://www.styra.com/blog/linting-rego-with-rego/) to determine whether
policies conform to Regal's linter rules.

## Goals

- Fast to traverse and process in Rego
- Usable without having to deal with quirks and inconsistencies
- As easy to read as the original AST JSON format

While this module provides a way to encode an `ast.Module` to an optimized JSON format, it does not provide a decoder.
In other words, there's no way to turn optimized AST JSON back into an `ast.Module` (or other AST types). While this
would be possible to do, there's no real need for that given our current use-case for this format, which is to help work
with the AST efficiently in Rego. Roast should not be considered a general purpose format for serializing the Rego AST.

## Differences

The following section outlines the differences between the original AST JSON format and the Roast format.

### Compact `location` format

The perhaps most visually apparent change to the AST JSON format is how `location` attributes are represented. These
attributes are **everywhere** in the AST, so optimizing these for fast traversal has a huge impact on both the size of
the format and the speed at which it can be processed.

In the original AST, a `location` is represented as an object:

```json
{
  "file": "p.rego",
  "row": 5,
  "col": 1,
  "text": "Y29sbGVjdGlvbg=="
}
```

And in the optimized format as a string:

```json
"5:1:Y29sbGVjdGlvbg=="
```

While this may come with a small cost for when the `location` is actually needed, it's a huge win for when it's not.
Having to `split` the result and parse the row and column values when needed occurs some overhead, but only a small
percentage of `location` attributes are commonly used in practice.

Note that the `file` attribute is omitted entirely in the optimized format, as this would otherwise have to be repeated
for each `location` value. This can easily be retrieved by other means.

### "Empty" rule and `else` bodies

Rego rules don't necessarily have a body, or at least not one that's printed. Examples of this include:

```rego
package policy

default rule := "value"

map["key"] := "value"

collection contains "value"
```

OPA represents such rules internally (that is, in the AST) as having a body with a single expression containing the
boolean value `true`. This creates a uniform way to represent rules, so a rule like:

```rego
collection contains "value"
```

Would in the AST be identical to:

```rego
collection contains "value" if {
    true
}
```

And in the OPA JSON AST format:

```json
{
  "body": [
    {
      "index": 0,
      "location": {
        "file": "p.rego",
        "row": 5,
        "col": 1,
        "text": "Y29sbGVjdGlvbg=="
      },
      "terms": {
        "location": {
          "file": "p.rego",
          "row": 5,
          "col": 1,
          "text": "Y29sbGVjdGlvbg=="
        },
        "type": "boolean",
        "value": true
      }
    }
  ],
  "head": {
    "name": "collection",
    "key": {
      "location": {
        "file": "p.rego",
        "row": 5,
        "col": 21,
        "text": "InZhbHVlIg=="
      },
      "type": "string",
      "value": "value"
    },
    "ref": [
      {
        "location": {
          "file": "p.rego",
          "row": 5,
          "col": 1,
          "text": "Y29sbGVjdGlvbg=="
        },
        "type": "var",
        "value": "collection"
      }
    ],
    "location": {
      "file": "p.rego",
      "row": 5,
      "col": 1,
      "text": "Y29sbGVjdGlvbiBjb250YWlucyAidmFsdWUi"
    }
  },
  "location": {
    "file": "p.rego",
    "row": 5,
    "col": 1,
    "text": "Y29sbGVjdGlvbg=="
  }
}
```

Notice how there's 20 lines of JSON just to represent the body, even though there isn't really one!


The optimized Rego AST format discards generated bodies entirely, and the same rule would be represented as:

```json
{
  "head": {
    "location": "5:1:Y29sbGVjdGlvbiBjb250YWlucyAidmFsdWUi",
    "name": "collection",
    "ref": [
      {
        "location": "5:1:Y29sbGVjdGlvbg==",
        "type": "var",
        "value": "collection"
      }
    ],
    "key": {
      "type": "string",
      "value": "value",
      "location": "5:21:InZhbHVlIg=="
    }
  },
  "location": "5:1:Y29sbGVjdGlvbg=="
}
```

Note that this applies equally to empty `else` bodies, which are represented the same way in the original AST, and
omitted entirely in the optimized format.

### Removed `annotations` attribute from module

OPA already attaches `annotations` to rules. With the Roast format attaching `package` and `subpackages` scoped
`annotations` to the `package` as well, there is no need to store `annotations` at the module level, as that's
effectively just duplicating data. Having this removed can save a considerable amount of space in well-documented
policies, as they should be!

### Removed `index` attribute from body expressions

In the original AST, each expression in a body carries a numeric `index` attribute. While this doesn't take much space,
it is largely redundant, as the same number can be inferred from the order of the expressions in the body array. It's
therefore been removed from the Roast format.

### Removed`name` attribute from rule heads

The `name` attribute found in the OPA AST for `rules` is unreliable, as it's not always present. The `ref`
attribute however always is.  While this doesn't come with any real cost in terms of AST size or performance, consistency
is key.

### Fixed inconsistencies in the original Rego AST

A few inconsistencies exist in the original AST JSON format:

- `comments` attributes having a `Text` attribute rather than the expected `text`
- `comments` attributes having a `Location` attribute rather than the expected `location`

Fixing these in the original format would be a breaking change. The Roast format corrects these inconsistencies, and
uses `text` and `location` consistently.

## Performance

While the numbers may vary some, the Roast format is currently about 40-50% smaller in size than the original AST JSON
format, and can be processed (in Rego, using `walk` and so on) about 1.25 times faster.

## Potential improvements

### Replace `text` in `location` string with end location

While it's not known how much of an impact it has on traversal speed / performance, a large chunk of the bytes in the
current optimized AST are base64-encoded `text` values in `location` strings. These could be replaced with the end
location of any given node. Rather than base64-encoding the the bytes of the text in serialization, we could instead
count the number of newlines in the text, and from that plus the number of bytes on the last line (if more than one)
determine the end location. It would then be assumed that the client has the means to translate that into the
equivalence of `text` where necessary. Regal could for example easily do this from `input.regal.file.lines`.

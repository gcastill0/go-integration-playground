## Details

The document present an array of three objects with two syntax errors in the document structure. 

* After the data array opening with `[` symbol, the first object is missing a `{` before the first `"id"` key. 

    The snippet shown below shows the first object in the original file - [See lines 1–24 of jdoc_original.json](https://github.com/gcastill0/go-integration-playground/blob/main/tasks/task1-jsonfix/jdoc_original.json#L1-24). Note specifically that the first object is missing the opening handlebar ('{') which should be located in line 2.

    ```json
    01  [
    02    
    03      "id": 1,
    04      "name": "Leanne Graham",
    05      "username": "Bret",
    06      "email": "Sincere@april.biz",
    07      "address": {
    08          "street": "Kulas Light",
    09          "suite": "Apt. 556",
    10          "city": "Gwenborough",
    11          "zipcode": "92998-3874",
    12          "geo": {
    13              "lat": "-37.3159",
    14              "lng": "81.1496"
    15          }
    16      },
    17      "phone": "1-770-736-8031 x56442",
    18      "website": "hildegard.org",
    19      "company": {
    20          "name": "Romaguera-Crona",
    21          "catchPhrase": "Multi-layered client-server neural-net",
    22            "bs": "harness real-time e-markets"
    23        }
    24    },
        ...
    ```

<br>

* The data structure never closes the array with a `]` symbol after the third object; the file moves on to the final `}`.

    The snippet shown below shows the first object in the original file - [See lines 1–24 of jdoc_original.json](https://github.com/gcastill0/go-integration-playground/blob/main/tasks/task1-jsonfix/jdoc_original.json#L1-24). Note specifically that the first object is missing the opening handlebar ('{') which should be located in line 2.

    The following snippet shows the last object in the dictionrary from original file - [See lines 48–71 of jdoc_original.json](https://github.com/gcastill0/go-integration-playground/blob/main/tasks/task1-jsonfix/jdoc_original.json#L1-24). Note specifically that the array is missing the closing bracker (']') which should be located in line 71.

    ```json
        ...
    48    {
    49    "id": 3,
    50   "name": "Clementine Bauch",
    51    "username": "Samantha",
    52    "email": "Nathan@yesenia.net",
    53    "address": {
    54      "street": "Douglas Extension",
    55        "suite": "Suite 847",
    56        "city": "McKenziehaven",
    57        "zipcode": "59590-4157",
    58        "geo": {
    59            "lat": "-68.6102",
    60            "lng": "-47.0653"
    61        }
    62    },
    63    "phone": "1-463-123-4447",
    64    "website": "ramiro.info",
    65    "company": {
    66        "name": "Romaguera-Jacobson",
    67        "catchPhrase": "Face to face bifurcated interface",
    68        "bs": "e-enable strategic applications"
    69    }
    70  }
    71
    ```

<br>

## Fixes

There are two the minimal fixes: 

1. insert a left handle bar `{` on line 2, right after the opening `[`, and 

2. add the closing right bracker, `]`, at the end of the array on line 71.

<br>

## JSON Grammar

With JSON, there two primary composite components: objects and arrays.

An **object** is an unordered collection of key–value pairs. The grammar requires that an object begin with a left brace, or handlebar, and end with a right brace, and that each member be a string key followed by a colon and then a value, with members separated by commas. Keys must be double-quoted strings; bare identifiers are not allowed. Semantically, an object corresponds to a dictionary or hash map from string to value, so lookups are by key, and the textual order of members carries no meaning in the data model.

An **array** is an ordered sequence of values. The grammar requires that an array begin with a left bracket and end with a right bracket, and that commas separate elements. Unlike objects, order in an array is semantically significant and corresponds to positional indices starting at zero, as expected from a dynamic array or vector. Arrays may be empty, may hold heterogeneous values, and may contain nested arrays or objects.

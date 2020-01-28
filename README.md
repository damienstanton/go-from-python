# Using Go from Python

This repository is a toy example of using Cgo and Python's C FFI to effectively reuse Go code
from Python. Tested on Go 1.12+, Python 3.8+. 

Notes:
- In the `lib` dir, one does not need the generated C header file. It can be useful, however,
for understanding how Cgo generates code.
- Required Python packages: To run tests, `pytest` is required. To try the example
notebook, `jupyter` is required.

## Usage

Build shared C lib, run Go and Python unit tests:
```sh
$ make
```

Cleanup:
```sh
$ make clean
```

Explore with Jupyter Notebook:
```sh
$ jupyter notebook # then open example.ipynb
```

© 2020 Damien Stanton

See LICENSE for details.

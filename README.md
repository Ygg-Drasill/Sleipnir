# Sleipnir

![image](https://github.com/Ygg-Drasill/Sleipnir/assets/151849979/9c9e9985-7e31-423e-a2e0-1d9a82d15291)

The Sleipnir compiler compiles Yggdrasill code to WebAssembly or JavaScript

# Testing
The main integration test uses samples located inside `test/samples/valid` and `test/samples/valid`.
All invalid samples are expected to produce an error when compiled, while valid samples are expected to compile without errors.
Optionally, for much higher test quality a result value can be expected from valid samples using the following format in the sample files:
```
<source code>
expect:<result>
```
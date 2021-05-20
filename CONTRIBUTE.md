# Docs To Contribute

## Tips For Testing

Since testing the `tty` input is difficult, you can mock the `gkp.GetKeyPushed()` behavior by changing the global variables such as `gkp.DummyMsg` and `gkp.DummyErr`.

If you set any string to `gkp.DummyMsg` then `gkp.GetKeyPushed()` will return that string instead of `tty` input.

```go
func TestGetKeyPushed_DummyMsg(t *testing.T) {
    expect := "a"

    gkp.DummyMsg = expect

    defer func() { gkp.DummyMsg = "" }()

    actual, err := gkp.GetKeyPushed()
    assert.Equal(t, expect, actual)
    assert.Nil(t, err)
}
```

If you set an `error` to `gkp.DummyErr` then `gkp.GetKeyPushed()` will return that error.

```go
func TestGetKeyPushed_DummyErr(t *testing.T) {
    expect := "dummy error occurred"

    errDummy := errors.New(expect)

    gkp.DummyErr = errDummy

    defer func() { gkp.DummyErr = nil }()

    msg, err := gkp.GetKeyPushed()

    assert.EqualErrorf(t, err, expect, "error message %s", "mal-formatted")
}
```

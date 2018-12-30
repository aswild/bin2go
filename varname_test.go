package bin2go

import "testing"

func TestCheckVarname(t *testing.T) {
    tests := []struct {
        f string // filename
        r bool   // expected return value
    }{
        {"helloWorld", true},
        {"hello_world", true},
        {"", false},
        {"_", false},
        {"0Hello", false},
        {"hello world", false},
        {"hello-world", false},
    }

    for _, c := range tests {
        r := CheckVarname(c.f)
        if r != c.r {
            t.Errorf("CheckVarname(%q) expected %t got %t", c.f, c.r, r)
        }
    }
}

func TestFilenameToVarname(t *testing.T) {
    tests := []struct {
        f string // filename
        v string // expected varname
        e bool   // expected error
    }{
        {"", "", true},
        {"_", "", true},
        {"example.conf", "example_conf", false},
        {"hello - world.txt", "hello_world_txt", false},
        {"0foo_____", "_0foo", false},
        {"Default.conf", "Default_conf", false},
    }

    for _, c := range tests {
        v, err := FilenameToVarname(c.f)
        if c.e && err == nil {
            t.Errorf("FilenameToVarname(%q) expected error but got success result %q", c.f, v)
        } else if !c.e && err != nil {
            t.Errorf("FilenameToVarname(%q) unexpected error: %v", c.f, err)
        } else if v != c.v {
            t.Errorf("FilenameToVarname(%q) expected %q got %q", c.f, c.v, v)
        }
    }
}

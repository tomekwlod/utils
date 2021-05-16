package env

// go test github.com/tomekwlod/utils/env -v -cover

import (
	"os"
	"testing"
)

const TESTENVNAME = "TESTENV"

func TestEnv(t *testing.T) {

	os.Setenv(TESTENVNAME, "")

	shouldBe := "ishouldbeused"

	res := Env(TESTENVNAME, shouldBe)

	if res != shouldBe {
		t.Errorf("Env() function returned `%s`, `%s` expected", res, shouldBe)
	} else {
		t.Logf("Env() default value case succeeded")
	}

	//

	os.Setenv(TESTENVNAME, shouldBe)

	res = Env(TESTENVNAME, "ishouldnotbeused")

	if res != shouldBe {
		t.Errorf("Env() function returned `%s`, `%s` expected", res, shouldBe)
	} else {
		t.Logf("Env() simple case succeeded")
	}

	//

	shouldBe = ""
	os.Setenv(TESTENVNAME, shouldBe)

	res = Env(TESTENVNAME, shouldBe)

	if res != shouldBe {
		t.Errorf("Env() function returned `%s`, `%s` expected", res, shouldBe)
	} else {
		t.Logf("Env() empty case succeeded")
	}
}

func TestIntEnv(t *testing.T) {

	os.Setenv(TESTENVNAME, "1")

	shouldBe := 1

	res := EnvInt(TESTENVNAME, 2)

	if res != shouldBe {
		t.Errorf("Env() function returned `%d`, `%d` expected", res, shouldBe)
	} else {
		t.Logf("Env() simple value case succeeded")
	}

	//

	os.Setenv(TESTENVNAME, "one")

	shouldBe = 1

	res = EnvInt(TESTENVNAME, 1)

	if res != shouldBe {
		t.Errorf("Env() function returned `%d`, `%d` expected", res, shouldBe)
	} else {
		t.Logf("Env() string value case succeeded")
	}

	//

	os.Setenv(TESTENVNAME, "")

	shouldBe = 1

	res = EnvInt(TESTENVNAME, 1)

	if res != shouldBe {
		t.Errorf("Env() function returned `%d`, `%d` expected", res, shouldBe)
	} else {
		t.Logf("Env() default value case succeeded")
	}
}

func TestIntBool(t *testing.T) {

	os.Setenv(TESTENVNAME, "true")

	shouldBe := true

	res := EnvBool(TESTENVNAME)

	if res != shouldBe {
		t.Errorf("Env() function returned `%t`, `%t` expected", res, shouldBe)
	} else {
		t.Logf("Env() simple value case succeeded")
	}

	//

	os.Setenv(TESTENVNAME, "t")

	shouldBe = true

	res = EnvBool(TESTENVNAME)

	if res != shouldBe {
		t.Errorf("Env() function returned `%t`, `%t` expected", res, shouldBe)
	} else {
		t.Logf("Env() `t` value case succeeded")
	}

	//

	os.Setenv(TESTENVNAME, "y")

	shouldBe = false

	res = EnvBool(TESTENVNAME)

	if res != shouldBe {
		t.Errorf("Env() function returned `%t`, `%t` expected", res, shouldBe)
	} else {
		t.Logf("Env() `y` value case succeeded")
	}

	//

	os.Setenv(TESTENVNAME, "FALSE")

	shouldBe = false

	res = EnvBool(TESTENVNAME)

	if res != shouldBe {
		t.Errorf("Env() function returned `%t`, `%t` expected", res, shouldBe)
	} else {
		t.Logf("Env() `FALSE` value case succeeded")
	}

	//

	os.Setenv(TESTENVNAME, "f")

	shouldBe = false

	res = EnvBool(TESTENVNAME)

	if res != shouldBe {
		t.Errorf("Env() function returned `%t`, `%t` expected", res, shouldBe)
	} else {
		t.Logf("Env() `f` value case succeeded")
	}

}

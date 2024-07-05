package utils

import "testing"

func TestConvertNameToRoute(t *testing.T) {

    test := "ProfilPublique"
    expected := "/profil/publique"
    actual := ConvertNameToRoute(test)

    if expected != actual {
        t.Fatalf("exected %s, got %s", expected, actual)
    }

    test = "Profil"
    expected = "/profil"
    actual = ConvertNameToRoute(test)

    if expected != actual {
        t.Fatalf("exected %s, got %s", expected, actual)
    }

}
func TestConvertNameToRouteName(t *testing.T) {

    test := "ProfilPublique"
    expected := "app_profil_publique"
    actual := ConvertNameToRouteName(test)

    if expected != actual {
        t.Fatalf("exected %s, got %s", expected, actual)
    }

    test = "Profil"
    expected  = "app_profil"
    actual = ConvertNameToRouteName(test)

    if expected != actual {
        t.Fatalf("exected %s, got %s", expected, actual)
    }

}

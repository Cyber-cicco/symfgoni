<?php

namespace {{.NameSpace}};
{{range .Imports}} {{.}} {{end}}

class {{.ClassName}} {

    public function __construct(
        {{range .Fields}}
        {{.}}
        {{end}}
    )
    {

    }
}

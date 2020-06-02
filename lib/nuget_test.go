package lib

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPatchNuget(t *testing.T) {
	xmlString := `<?xml version="1.0" encoding="utf-8"?>
<configuration>
	<packageSources>
		<add key="MySource"/>
	</packageSources>
</configuration>
`

	reader := strings.NewReader(xmlString)
	buffer := new(bytes.Buffer)

	err := PatchNuget(reader, buffer, "MySource", "username", "password")
	require.NoError(t, err, "Could not patch nuget")

	expectedString := `<?xml version="1.0" encoding="utf-8"?>
<configuration>
	<packageSources>
		<add key="MySource"/>
	</packageSources>
	<packageSourceCredentials>
		<MySource>
			<add key="Username" value="username"/>
			<add key="ClearTextPassword" value="password"/>
		</MySource>
	</packageSourceCredentials>
</configuration>
`

	assert.EqualValues(t, expectedString, buffer.String(), "XML must be equal")
}

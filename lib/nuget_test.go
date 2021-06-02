package lib

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPatchNuget_WithoutPreviousPackageSourceCredentials(t *testing.T) {
	source := faker.Word()
	username := faker.Username()
	password := faker.Password()

	xmlString := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<configuration>
	<packageSources>
		<add key="%s"/>
	</packageSources>
</configuration>
`, source)

	reader := strings.NewReader(xmlString)
	buffer := new(bytes.Buffer)

	err := PatchNuget(reader, buffer, source, username, password)
	require.NoError(t, err, "Could not patch nuget")

	expectedString := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<configuration>
	<packageSources>
		<add key="%[1]s"/>
	</packageSources>
	<packageSourceCredentials>
		<%[1]s>
			<add key="Username" value="%[2]s"/>
			<add key="ClearTextPassword" value="%[3]s"/>
		</%[1]s>
	</packageSourceCredentials>
</configuration>
`, source, username, password)

	assert.EqualValues(t, expectedString, buffer.String(), "XML must be equal")
}

func TestPatchNuget_WithPreviousPackageSourceCredentials(t *testing.T) {
	source1 := faker.Word()
	username1 := faker.Username()
	password1 := faker.Password()

	source2 := faker.Word()
	wrong_username2 := faker.Username()
	wrong_password2 := faker.Password()
	username2 := faker.Username()
	password2 := faker.Password()

	xmlString := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<configuration>
	<packageSources>
		<add key="%[1]s"/>
		<add key="%[4]s"/>
	</packageSources>
	<packageSourceCredentials>
		<%[1]s>
			<add key="Username" value="%[2]s"/>
			<add key="ClearTextPassword" value="%[3]s"/>
		</%[1]s>
		<%[4]s>
			<add key="Username" value="%[5]s"/>
			<add key="ClearTextPassword" value="%[6]s"/>
		</%[4]s>
	</packageSourceCredentials>
</configuration>
`,
		source1, username1, password1,
		source2, wrong_username2, wrong_password2)

	reader := strings.NewReader(xmlString)
	buffer := new(bytes.Buffer)

	err := PatchNuget(reader, buffer, source2, username2, password2)
	require.NoError(t, err, "Could not patch nuget")

	expectedString := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<configuration>
	<packageSources>
		<add key="%[1]s"/>
		<add key="%[4]s"/>
	</packageSources>
	<packageSourceCredentials>
		<%[1]s>
			<add key="Username" value="%[2]s"/>
			<add key="ClearTextPassword" value="%[3]s"/>
		</%[1]s>
		<%[4]s>
			<add key="Username" value="%[5]s"/>
			<add key="ClearTextPassword" value="%[6]s"/>
		</%[4]s>
	</packageSourceCredentials>
</configuration>
`,
		source1, username1, password1,
		source2, username2, password2)

	assert.EqualValues(t, expectedString, buffer.String(), "XML must be equal")
}

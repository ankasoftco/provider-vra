package models 
type RemoteAccessSpecification struct {

	// One of four authentication types.
	// `generatedPublicPrivateKey`: The provisioned machine generates the public/private key pair and enables SSH to use them without user input.
	// `publicPrivateKey`: The user enters the private key in the SSH command. See remoteAccess.sshKey.
	// `usernamePassword`: The user enters a username and password for remote access.
	// `keyPairName`: The user enters an already existing keyPair name. See remoteAccess.keyPair
	// Example: publicPrivateKey
	// Required: true
	Authentication *string `json:"authentication"`

	// Key Pair Name.
	KeyPair string `json:"keyPair,omitempty"`

	// Remote access password for the Azure machine.
	Password string `json:"password,omitempty"`

	// In key pair authentication, the public key on the provisioned machine. Users are expected to log in with their private key and a default username from the cloud provider. An AWS Ubuntu image comes with default user ubuntu, and Azure comes with default user azureuser. To log in by SSH:
	// `ssh -i <private-key-path> ubuntu@52.90.80.153`
	// `ssh -i <private-key-path> azureuser@40.76.14.255`
	// Example: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCu74dLkAGGYIgNuszEAM0HaS2Y6boTPw+HqsFmtPSOpxPQoosws/OWGZlW1uue6Y4lIvdRqZOaLK+2di5512etY67ZwFHc5h1kx4az433DsnoZhIzXEKKI+EXfH/w72CIyG/uVhIzmA4FvRVQKXinE1vaVen6v1CBQEZibx9RXrVRP1VRibsKFRXYxywNEl1VtPK7KaxCEYO9IXi4SKVulSAhOVequwjlo5E8bKNT61/g/YyMvwCbaTTPPeCpS/7i+JHYY3QZ8fQY/Syn+bOFpKCCHl+7VpsL8gjWe6fI4bUp6KUiW7ZkQpL/47rxawKnRMKKEU9P0ICp3RRB39lXT
	SSHKey string `json:"sshKey,omitempty"`

	// Remote access username for the Azure machine.
	Username string `json:"username,omitempty"`
}


package models 
type MachineBootConfig struct {

	// A valid cloud config data in json-escaped yaml syntax
	// Example: #cloud-config\nrepo_update: true\nrepo_upgrade: all\n\npackages:\n - mysql-server\n\nruncmd:\n - sed -e '/bind-address/ s/^#*/#/' -i /etc/mysql/mysql.conf.d/mysqld.cnf\n - service mysql restart\n - mysql -e \"GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'mysqlpassword';\"\n - mysql -e \"FLUSH PRIVILEGES;\"\n
	Content string `json:"content,omitempty"`
}


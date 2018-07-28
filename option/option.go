package option

import "flag"

var (
	Opt *Options
)

// Options model for commandline arguments
type Options struct {
	HostName  string
	Port      string
	UserName  string
	Password  string
	Databases string

	MySQLDumpPath string
	BackupPath    string

	Security string

	// DatabaseRowCountTreshold int
	// TableRowCountTreshold    int
	// BatchSize                int
	// ForceSplit               bool

	// AdditionalMySQLDumpArgs string

	// Verbosity int

	// OutputDirectory        string
	// DefaultsProvidedByUser bool
	// ExecutionStartDate     time.Time

	// DailyRotation   int
	// WeeklyRotation  int
	// MonthlyRotation int
}

func InitOpt() {
	Opt = &Options{}
	flag.StringVar(&Opt.HostName, "hostname", "localhost", "Hostname of the mysql server to connect to")
	flag.StringVar(&Opt.Port, "port", "3306", "Port of the mysql server to connect to")
	flag.StringVar(&Opt.UserName, "username", "root", "username of the mysql server to connect to")
	flag.StringVar(&Opt.Password, "password", "root", "password of the mysql server to connect to")
	flag.StringVar(&Opt.Databases, "databases", "--all-databases", "List of databases as comma seperated values to dump. OBS: If not specified, --all-databases is the default")
	flag.StringVar(&Opt.MySQLDumpPath, "mysqldump-path", "/usr/bin/mysqldump", "Absolute path for mysqldump executable.")
	flag.StringVar(&Opt.BackupPath, "backup-path", "", "Default is the value of os.Getwd(). The backup files will be placed to output-dir /{DATABASE_NAME}/{DATABASE_NAME}_{TABLENAME|SCHEMA|DATA|ALL}_{TIMESTAMP}.sql")
	flag.StringVar(&Opt.Security, "en-passwd", "", "Encrypt password")
	flag.Parse()
}

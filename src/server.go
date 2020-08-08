//

package GroupWebApp

type ServerAction int

const (
	Auto            ServerAction = 0
	Run             ServerAction = 1
	Setup                        = 2
	RebuildDatabase              = 3
	PasswordReset                = 4
)

func Serve(action ServerAction, configFile *string) {

	switch action {
	case Run:
		// Just run. If server is not configured, then error.
		break
	case Setup:
		// Just serve the "setup" pages.
		break
	case BuildDatabase:
		// Just build/recreate all database tables (if necessary, delete all data). If server database is not configured, then error.
		break
	case PasswordReset:
		// Just reset all admin passwords in database and output them to the user. If server is not configured, then error.
		break
	default:
		// Auto: serve the "setup" pages if server is not configured, or just run if server is configured.
	}

}

//

package GroupWebApp

import "database/sql"

const SQL_AUTH_ = ""

const SQL_USER_GETBYID = string{"SELECT id, ", ""}


type DatabaseConnection connection sql.DB

func InitDB()
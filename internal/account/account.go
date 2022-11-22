package account

import (
	"context"
	"fmt"
	"github.com/aasumitro/posbe/internal/account/repository/sql"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/aasumitro/posbe/pkg/sql/query"
	"github.com/gin-gonic/gin"
)

func InitAccountModule(config *config.Config, router *gin.Engine) {
	selectAll := query.SQLSelectBuilder{}.
		ForTable("roles").
		Build()

	selectAllWhere := query.SQLSelectBuilder{}.
		ForTable("roles").
		AddWhere([]string{"id = 1", "name = 'lorem'"}).
		Build()

	selectSpecifiedWhere := query.SQLSelectBuilder{}.
		ForTable("roles").
		JustField("id, name").
		AddWhere([]string{"id = 1"}).
		Build()

	fmt.Println(selectAll)
	fmt.Println(selectAllWhere)
	fmt.Println(selectSpecifiedWhere)

	ctx := context.Background()
	roleRepository := sql.NewRoleSQlRepository(config.GetDbConn())

	roles, err := roleRepository.All(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(roles)

	role, err := roleRepository.Find(ctx, []string{"id = 1"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(role)
}

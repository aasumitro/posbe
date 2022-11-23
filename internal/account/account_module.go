package account

import (
	"context"
	"fmt"
	"github.com/aasumitro/posbe/internal/account/repository/sql"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/aasumitro/posbe/pkg/sql/query"
	"github.com/gin-gonic/gin"
)

func InitAccountModule(ctx context.Context, config *config.Config, router *gin.Engine) {
	selectAll := query.SQLSelectBuilder{}.
		Table("roles").
		Build()

	selectAllWhere := query.SQLSelectBuilder{}.
		Table("roles").
		Where("id = 1").
		Build()

	q := query.SQLSelectBuilder{}.
		Table("lorem").
		Field("ipsum as i").
		HasLimit(2).
		HasOffset(3).
		GroupBy("lorem").
		OrderBy("ipsum").
		Join("mamun ON lorem.id = mamun.lorem_id").
		InnerJoin("mamun ON lorem.id = mamun.lorem_id").
		CrossJoin("mamun ON lorem.id = mamun.lorem_id").
		LeftOuterJoin("mamun ON lorem.id = mamun.lorem_id").
		RightOuterJoin("mamun ON lorem.id = mamun.lorem_id").
		FullOuterJoin("mamun ON lorem.id = mamun.lorem_id").
		Where("").
		Build()

	fmt.Println(selectAll)
	fmt.Println(selectAllWhere)
	fmt.Println(q)

	roleRepository := sql.NewRoleSQlRepository(config.GetDbConn())

	roles, err := roleRepository.All(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(roles)

	role, err := roleRepository.Find(ctx, "id = 1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(role)
}

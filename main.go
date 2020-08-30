package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

    router.GET("/", func(ctx *gin.Context){
        db, err := sql.Open("mysql", "root:password@tcp(172.20.0.2:3306)/sample_docker_compose")
        if err != nil {
            panic(err.Error())
        }
        defer db.Close()

        rows, err := db.Query("SELECT * FROM user")
        if err != nil {
            panic(err.Error())
        }

        columns, err := rows.Columns()
        if err != nil {
            panic(err.Error())
        }

        values := make([]sql.RawBytes, len(columns))
        scanArgs := make([]interface{}, len(values))
        for i := range values {
            scanArgs[i] = &values[i]
        }

        var data []string
        for rows.Next() {
            err = rows.Scan(scanArgs...)
            if err != nil {
                panic(err.Error())
            }

            var value string
            for i, col := range values {
                if col == nil {
                    value = "NULL"
                } else {
                    value = string(col)
                }
                data = append(data, columns[i] + ": " + value)
            }
        }

        ctx.HTML(200, "index.html", gin.H{"data":data})
    })

    router.Run()
}

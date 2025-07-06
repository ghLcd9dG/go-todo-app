package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time" // Import the time package
)

const dataFile = "todos.json"

// Todo 结构体表示一个待办事项
type Todo struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"` // Add this field
}

// loadTodos 从文件中加载待办事项
func loadTodos() ([]Todo, error) {
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil // 文件不存在，返回空列表
		}
		return nil, fmt.Errorf("无法读取文件: %w", err)
	}

	var todos []Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, fmt.Errorf("无法解析 JSON: %w", err)
	}
	return todos, nil
}

// saveTodos 将待办事项保存到文件中
func saveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("无法编码 JSON: %w", err)
	}
	if err := ioutil.WriteFile(dataFile, data, 0644); err != nil {
		return fmt.Errorf("无法写入文件: %w", err)
	}
	return nil
}

// addTodo 添加一个新的待办事项
func addTodo(task string, todos *[]Todo) {
	newID := 1
	if len(*todos) > 0 {
		newID = (*todos)[len(*todos)-1].ID + 1
	}

	*todos = append(*todos, Todo{
		ID:        newID,
		Task:      task,
		Completed: false,
		CreatedAt: time.Now().UTC(), // Set creation time to UTC
	})
	fmt.Printf("已添加待办事项: \"%s\" (ID: %d)\n", task, newID)
}

// listTodos 列出所有待办事项
func listTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("待办事项列表为空。")
		return
	}
	fmt.Println("--- 待办事项列表 ---")
	for _, todo := range todos {
		status := " "
		if todo.Completed {
			status = "x"
		}
		// Display the creation date (formatted)
		fmt.Printf("[%s] %d. %s (创建时间: %s UTC)\n", status, todo.ID, todo.Task, todo.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	fmt.Println("--------------------")
}

// completeTodo 将指定ID的待办事项标记为完成
func completeTodo(id int, todos *[]Todo) {
	found := false
	for i := range *todos {
		if (*todos)[i].ID == id {
			(*todos)[i].Completed = true
			found = true
			fmt.Printf("已将待办事项 ID %d 标记为完成。\n", id)
			break
		}
	}
	if !found {
		fmt.Printf("未找到 ID 为 %d 的待办事项。\n", id)
	}
}

// printUsage 打印使用说明
func printUsage() {
	fmt.Println("\n使用方法:")
	fmt.Println("  add <任务描述>    - 添加一个新的待办事项")
	fmt.Println("  list              - 列出所有待办事项")
	fmt.Println("  complete <ID>     - 将指定 ID 的待办事项标记为完成")
	fmt.Println("  help              - 显示此帮助信息")
	fmt.Println("  exit              - 退出程序")
}

func main() {
	todos, err := loadTodos()
	if err != nil {
		fmt.Printf("加载待办事项失败: %s\n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	printUsage()

	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("请提供待办事项的描述。")
				continue
			}
			task := strings.Join(parts[1:], " ")
			addTodo(task, &todos)
			if err := saveTodos(todos); err != nil {
				fmt.Printf("保存待办事项失败: %s\n", err)
			}
		case "list":
			listTodos(todos)
		case "complete":
			if len(parts) < 2 {
				fmt.Println("请提供要完成的待办事项的 ID。")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("无效的 ID。请提供一个数字。")
				continue
			}
			completeTodo(id, &todos)
			if err := saveTodos(todos); err != nil {
				fmt.Printf("保存待办事项失败: %s\n", err)
			}
		case "help":
			printUsage()
		case "exit":
			fmt.Println("退出待办事项列表程序。")
			return
		default:
			fmt.Println("未知命令。输入 'help' 查看可用命令。")
		}
	}
}

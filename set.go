package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ==================== КРАСНО-ЧЕРНОЕ ДЕРЕВО ====================

type Color int

const (
	RED Color = iota
	BLACK
)

type RBNode struct {
	data   string
	color  Color
	parent *RBNode
	left   *RBNode
	right  *RBNode
}

func NewRBNode(value string) *RBNode {
	return &RBNode{
		data:   value,
		color:  RED,
		parent: nil,
		left:   nil,
		right:  nil,
	}
}

type RBTree struct {
	root *RBNode
}

func NewRBTree() *RBTree {
	return &RBTree{root: nil}
}

// ==================== ФУНКЦИИ КРАСНО-ЧЕРНОГО ДЕРЕВА ====================

func leftRotate(tree *RBTree, x *RBNode) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}
	y.left = x
	x.parent = y
}

func rightRotate(tree *RBTree, x *RBNode) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}
	y.right = x
	x.parent = y
}

func fixAdd(tree *RBTree, z *RBNode) {
	for z.parent != nil && z.parent.color == RED {
		if z.parent.parent == nil {
			break
		}

		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right

			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					leftRotate(tree, z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				rightRotate(tree, z.parent.parent)
			}
		} else {
			y := z.parent.parent.left

			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					rightRotate(tree, z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				leftRotate(tree, z.parent.parent)
			}
		}

		if z == tree.root {
			break
		}
	}
	tree.root.color = BLACK
}

func AddRBNode(tree *RBTree, value string) {
	newNode := NewRBNode(value)
	if tree.root == nil {
		newNode.color = BLACK
		tree.root = newNode
		return
	}
	address := tree.root
	for {
		if value < address.data {
			if address.left == nil {
				newNode.parent = address
				address.left = newNode
				fixAdd(tree, newNode)
				return
			} else {
				address = address.left
			}
		} else {
			if address.right == nil {
				newNode.parent = address
				address.right = newNode
				fixAdd(tree, newNode)
				return
			} else {
				address = address.right
			}
		}
	}
}

func GetRBNode(tree *RBTree, value string) *RBNode {
	if tree.root == nil {
		return nil
	}

	address := tree.root
	for address != nil {
		if value == address.data {
			return address
		} else if value < address.data {
			address = address.left
		} else {
			address = address.right
		}
	}

	return nil
}

func treeMinimum(node *RBNode) *RBNode {
	if node == nil {
		return nil
	}
	address := node
	for address.left != nil {
		address = address.left
	}
	return address
}

func deleteFix(tree *RBTree, x *RBNode) {
	if x == nil {
		return
	}

	for x != tree.root && x.color == BLACK {
		if x == x.parent.left {
			w := x.parent.right
			if w == nil {
				break
			}

			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				leftRotate(tree, x.parent)
				w = x.parent.right
				if w == nil {
					break
				}
			}

			leftBlack := (w.left == nil) || (w.left.color == BLACK)
			rightBlack := (w.right == nil) || (w.right.color == BLACK)

			if leftBlack && rightBlack {
				w.color = RED
				x = x.parent
			} else {
				if w.right == nil || w.right.color == BLACK {
					if w.left != nil {
						w.left.color = BLACK
					}
					w.color = RED
					rightRotate(tree, w)
					w = x.parent.right
					if w == nil {
						break
					}
				}

				w.color = x.parent.color
				x.parent.color = BLACK
				if w.right != nil {
					w.right.color = BLACK
				}
				leftRotate(tree, x.parent)
				x = tree.root
			}
		} else {
			w := x.parent.left
			if w == nil {
				break
			}

			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				rightRotate(tree, x.parent)
				w = x.parent.left
				if w == nil {
					break
				}
			}

			leftBlack := (w.left == nil) || (w.left.color == BLACK)
			rightBlack := (w.right == nil) || (w.right.color == BLACK)

			if leftBlack && rightBlack {
				w.color = RED
				x = x.parent
			} else {
				if w.left == nil || w.left.color == BLACK {
					if w.right != nil {
						w.right.color = BLACK
					}
					w.color = RED
					leftRotate(tree, w)
					w = x.parent.left
					if w == nil {
						break
					}
				}

				w.color = x.parent.color
				x.parent.color = BLACK
				if w.left != nil {
					w.left.color = BLACK
				}
				rightRotate(tree, x.parent)
				x = tree.root
			}
		}
	}

	if x != nil {
		x.color = BLACK
	}
}

func transplant(tree *RBTree, u *RBNode, v *RBNode) {
	if u.parent == nil {
		tree.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

func DelRBNode(tree *RBTree, value string) {
	z := GetRBNode(tree, value)
	if z == nil {
		return
	}

	y := z
	var x *RBNode
	yOriginalColor := y.color

	if z.left == nil {
		x = z.right
		transplant(tree, z, x)
	} else if z.right == nil {
		x = z.left
		transplant(tree, z, x)
	} else {
		y = treeMinimum(z.right)
		yOriginalColor = y.color
		x = y.right

		if y.parent != z {
			transplant(tree, y, x)
			y.right = z.right
			if y.right != nil {
				y.right.parent = y
			}
		} else {
			if x != nil {
				x.parent = y
			}
		}

		transplant(tree, z, y)
		y.left = z.left
		if y.left != nil {
			y.left.parent = y
		}
		y.color = z.color
	}

	if yOriginalColor == BLACK {
		if x != nil {
			deleteFix(tree, x)
		}
	}
}

// ==================== МНОЖЕСТВО НА ОСНОВЕ КРАСНО-ЧЕРНОГО ДЕРЕВА ====================

type StringSet struct {
	tree *RBTree
}

func NewStringSet() *StringSet {
	return &StringSet{tree: NewRBTree()}
}

func (s *StringSet) Add(value string) {
	if s.Contains(value) {
		return
	}
	AddRBNode(s.tree, value)
}

func (s *StringSet) Remove(value string) {
	DelRBNode(s.tree, value)
}

func (s *StringSet) Contains(value string) bool {
	return GetRBNode(s.tree, value) != nil
}

func (s *StringSet) Size() int {
	count := 0
	s.ForEach(func(value string) {
		count++
	})
	return count
}

func (s *StringSet) Empty() bool {
	return s.tree.root == nil
}

func (s *StringSet) Clear() {
	s.tree.root = nil
}

func (s *StringSet) UnionWith(other *StringSet) *StringSet {
	result := NewStringSet()
	s.ForEach(func(value string) {
		result.Add(value)
	})
	other.ForEach(func(value string) {
		result.Add(value)
	})
	return result
}

func (s *StringSet) IntersectionWith(other *StringSet) *StringSet {
	result := NewStringSet()
	s.ForEach(func(value string) {
		if other.Contains(value) {
			result.Add(value)
		}
	})
	return result
}

func (s *StringSet) DifferenceWith(other *StringSet) *StringSet {
	result := NewStringSet()
	s.ForEach(func(value string) {
		if !other.Contains(value) {
			result.Add(value)
		}
	})
	return result
}

func (s *StringSet) ForEach(f func(string)) {
	inOrderTraversal(s.tree.root, f)
}

func inOrderTraversal(node *RBNode, f func(string)) {
	if node == nil {
		return
	}
	inOrderTraversal(node.left, f)
	f(node.data)
	inOrderTraversal(node.right, f)
}

func (s *StringSet) ToSlice() []string {
	var result []string
	s.ForEach(func(value string) {
		result = append(result, value)
	})
	return result
}

// ==================== БАЗА ДАННЫХ МНОЖЕСТВ ====================

type SetDatabase struct {
	sets map[string]*StringSet
}

func NewSetDatabase() *SetDatabase {
	return &SetDatabase{
		sets: make(map[string]*StringSet),
	}
}

func (db *SetDatabase) AddSet(name string, set *StringSet) {
	db.sets[name] = set
}

func (db *SetDatabase) RemoveSet(name string) {
	delete(db.sets, name)
}

func (db *SetDatabase) GetSet(name string) *StringSet {
	if set, exists := db.sets[name]; exists {
		return set
	}
	return nil
}

func (db *SetDatabase) ContainsSet(name string) bool {
	_, exists := db.sets[name]
	return exists
}

func (db *SetDatabase) ClearSet(name string) {
	if set, exists := db.sets[name]; exists {
		set.Clear()
	}
}

func (db *SetDatabase) GetSetNames() []string {
	names := make([]string, 0, len(db.sets))
	for name := range db.sets {
		names = append(names, name)
	}
	return names
}

func (db *SetDatabase) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	if _, err := writer.WriteString(fmt.Sprintf("SETS %d\n", len(db.sets))); err != nil {
		return err
	}

	for name, set := range db.sets {
		setData := set.ToSlice()

		if _, err := writer.WriteString(fmt.Sprintf("SET %s\n", name)); err != nil {
			return err
		}

		if _, err := writer.WriteString(fmt.Sprintf("%d\n", len(setData))); err != nil {
			return err
		}

		for _, data := range setData {
			if _, err := writer.WriteString(fmt.Sprintf("%s\n", data)); err != nil {
				return err
			}
		}
	}

	return writer.Flush()
}

func (db *SetDatabase) LoadFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		db.sets = make(map[string]*StringSet)
		return
	}
	defer file.Close()

	db.sets = make(map[string]*StringSet)
	scanner := bufio.NewScanner(file)

	var currentSetName string
	var elementsToRead int = -1
	var readingSet bool = false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if readingSet && elementsToRead > 0 {
			set := db.sets[currentSetName]
			if set == nil {
				set = NewStringSet()
				db.sets[currentSetName] = set
			}
			set.Add(line)
			elementsToRead--
			if elementsToRead == 0 {
				readingSet = false
				currentSetName = ""
			}
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "SET":
			if len(parts) >= 2 {
				currentSetName = parts[1]
				readingSet = true
				elementsToRead = -1
			}
		default:
			if readingSet && elementsToRead == -1 {
				if count, err := strconv.Atoi(line); err == nil {
					elementsToRead = count
					db.sets[currentSetName] = NewStringSet()
					if elementsToRead == 0 {
						readingSet = false
						currentSetName = ""
					}
				} else {
					readingSet = false
					currentSetName = ""
				}
			}
		}
	}
}

// ==================== ГЛАВНАЯ ФУНКЦИЯ ====================

func main() {
	if len(os.Args) < 5 {
		printUsage()
		return
	}

	var filename, query string

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch arg {
		case "--file":
			if i+1 < len(os.Args) {
				filename = os.Args[i+1]
				i++
			}
		case "--query":
			if i+1 < len(os.Args) {
				query = os.Args[i+1]
				i++
			}
		}
	}

	if filename == "" || query == "" {
		printUsage()
		return
	}

	database := NewSetDatabase()
	database.LoadFromFile(filename)

	tokens := strings.Fields(query)
	if len(tokens) == 0 {
		fmt.Println("Неверный формат запроса")
		return
	}

	operation := tokens[0]
	needSave := false

	switch {
	case operation == "SETCREATE" && len(tokens) == 2:
		setName := tokens[1]
		if !database.ContainsSet(setName) {
			database.AddSet(setName, NewStringSet())
			needSave = true
		} else {
			fmt.Printf("Множество '%s' уже существует\n", setName)
		}

	case operation == "SETADD" && len(tokens) == 3:
		setName := tokens[1]
		value := tokens[2]
		set := database.GetSet(setName)
		if set != nil {
			set.Add(value)
			needSave = true
		} else {
			fmt.Printf("Множество '%s' не найдено\n", setName)
		}

	case operation == "SETDEL" && len(tokens) == 3:
		setName := tokens[1]
		value := tokens[2]
		set := database.GetSet(setName)
		if set != nil {
			set.Remove(value)
			needSave = true
		} else {
			fmt.Printf("Множество '%s' не найдено\n", setName)
		}

	case operation == "SET_AT" && len(tokens) == 3:
		setName := tokens[1]
		value := tokens[2]
		set := database.GetSet(setName)
		if set != nil {
			if set.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		} else {
			fmt.Printf("Множество '%s' не найдено\n", setName)
		}

	case operation == "SET_UNION" && len(tokens) == 4:
		setName1 := tokens[1]
		setName2 := tokens[2]
		resultName := tokens[3]
		set1 := database.GetSet(setName1)
		set2 := database.GetSet(setName2)
		if set1 != nil && set2 != nil {
			result := set1.UnionWith(set2)
			database.AddSet(resultName, result)
			needSave = true
		} else {
			fmt.Println("Одно из множеств не найдено")
		}

	case operation == "SET_INTERSECT" && len(tokens) == 4:
		setName1 := tokens[1]
		setName2 := tokens[2]
		resultName := tokens[3]
		set1 := database.GetSet(setName1)
		set2 := database.GetSet(setName2)
		if set1 != nil && set2 != nil {
			result := set1.IntersectionWith(set2)
			database.AddSet(resultName, result)
			needSave = true
		} else {
			fmt.Println("Одно из множеств не найдено")
		}

	case operation == "SET_DIFF" && len(tokens) == 4:
		setName1 := tokens[1]
		setName2 := tokens[2]
		resultName := tokens[3]
		set1 := database.GetSet(setName1)
		set2 := database.GetSet(setName2)
		if set1 != nil && set2 != nil {
			result := set1.DifferenceWith(set2)
			database.AddSet(resultName, result)
			needSave = true
		} else {
			fmt.Println("Одно из множеств не найдено")
		}

	case operation == "SET_PRINT" && len(tokens) == 2:
		setName := tokens[1]
		set := database.GetSet(setName)
		if set != nil {
			if set.Empty() {
				fmt.Println("(пусто)")
			} else {
				set.ForEach(func(elem string) {
					fmt.Println(elem)
				})
			}
		} else {
			fmt.Printf("Множество '%s' не найдено\n", setName)
		}

	case operation == "SET_LIST":
		setNames := database.GetSetNames()
		if len(setNames) == 0 {
			fmt.Println("(нет множеств)")
		} else {
			for _, name := range setNames {
				set := database.GetSet(name)
				fmt.Printf("%s (%d элементов)\n", name, set.Size())
			}
		}

	case operation == "SETREMOVE" && len(tokens) == 2:
		setName := tokens[1]
		if database.ContainsSet(setName) {
			database.RemoveSet(setName)
			needSave = true
		} else {
			fmt.Printf("Множество '%s' не найдено\n", setName)
		}

	case operation == "SETCLEAR" && len(tokens) == 2:
		setName := tokens[1]
		set := database.GetSet(setName)
		if set != nil {
			set.Clear()
			needSave = true
		} else {
			fmt.Printf("Множество '%s' не найдено\n", setName)
		}

	default:
		fmt.Printf("Неизвестная команда или неверный формат: %s\n", operation)
		return
	}

	if needSave {
		if err := database.SaveToFile(filename); err != nil {
			fmt.Printf("Ошибка сохранения: %v\n", err)
		}
	}
}

func printUsage() {
	fmt.Println("Использование: program --file <filename> --query <query>")
	fmt.Println("Доступные операции:")
	fmt.Println("  SETCREATE <setname> - создать новое множество")
	fmt.Println("  SETADD <setname> <value> - добавить элемент в множество")
	fmt.Println("  SETDEL <setname> <value> - удалить элемент из множества")
	fmt.Println("  SET_AT <setname> <value> - проверить наличие элемента")
	fmt.Println("  SET_UNION <setname1> <setname2> <resultname> - объединение множеств")
	fmt.Println("  SET_INTERSECT <setname1> <setname2> <resultname> - пересечение множеств")
	fmt.Println("  SET_DIFF <setname1> <setname2> <resultname> - разность множеств")
	fmt.Println("  SET_PRINT <setname> - вывести все элементы множества")
	fmt.Println("  SET_LIST - вывести список всех множеств")
	fmt.Println("  SETREMOVE <setname> - удалить множество")
	fmt.Println("  SETCLEAR <setname> - очистить все элементы множества")
}

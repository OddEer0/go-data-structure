package rbtree

// Новый добавляемый элемент всегда является красным
// При добавлений красного баланс может быть нарушен 5 раз
// 1) Добавление корня красного цвета: Просто в конце балансировки всегда красим корень в черный
// 2) Если родитель был красным и (Не надо)дядя красный. Дед не корень: Красим дядю и родителя в черный, деда в красный. Дальше балансируем деда так как родитель деда может оказаться красным
// 3) Если родитель был красным и (Не надо)дядя красный. Дед корень: Делаем то же самое что и 2. Корень в конце балансировки всегда красится в черный
// 4) Если родитель был красным и (Не надо)дядя черный. При этом новый элемент добавляется зигзагом(родитель слева от дяди, новый будет справа от родителя и наобарот): Делаем левый или правый поворот относительно родителя и переносим состояние в 5й случай
// 5) Если родитель был красным и (Не надо)дядя черный. При этом новый элемент добавляется прямолинейно(родитель слева от дяди, новый будет слева от родителя и наобарот): Красим родителя в черный. Деда делаем красным и делаем правый или левый поворот относительно него

func (t *Tree[T, K]) balance(newNode *Node[T, K]) {
	var uncle *Node[T, K]
	for newNode.parent.color == red { // Убеждаемся что родитель красный
		if newNode.parent == newNode.parent.parent.left { // Убеждаемся, что родитель находится с левой стороны от деда
			uncle = newNode.parent.parent.right // Находим брата
			// Пункты 2 и 3
			if uncle.color == red {
				newNode.parent.color = black
				uncle.color = black
				newNode.parent.parent.color = red
				newNode = newNode.parent.parent
			} else { // Пункт 4 и 5
				// Пункт 4
				if newNode == newNode.parent.right {
					newNode = newNode.parent
					t.leftRotate(newNode)
				}
				// Пункт 5
				newNode.parent.color = black
				newNode.parent.parent.color = red
				t.rightRotate(newNode.parent.parent)
			}
		} else {
			uncle = newNode.parent.parent.left // Находим дядю
			// Пункты 2 и 3
			if uncle.color == red {
				newNode.parent.color = black
				uncle.color = black
				newNode.parent.parent.color = red
				newNode = newNode.parent.parent
			} else { // Пункт 4 и 5
				// Пункт 4
				if newNode == newNode.parent.left {
					newNode = newNode.parent
					t.rightRotate(newNode)
				}
				// Пункт 5
				newNode.parent.color = black
				newNode.parent.parent.color = red
				t.leftRotate(newNode.parent.parent)
			}
		}
	}

	// Пункт 1 и частично 3
	t.root.color = black
}

func (t *Tree[T, K]) leftRotate(node *Node[T, K]) {
	tmp := node.right
	node.right = tmp.left
	if node.right.NotNilNode() {
		node.right.parent = node
	}
	tmp.parent = node.parent
	if node == t.root {
		t.root = tmp
	} else if node == node.parent.left {
		node.parent.left = tmp
	} else {
		node.parent.right = tmp
	}
	tmp.left = node
	node.parent = tmp
}

func (t *Tree[T, K]) rightRotate(node *Node[T, K]) {
	tmp := node.left
	node.left = tmp.right
	if node.left.NotNilNode() {
		node.left.parent = node
	}
	tmp.parent = node.parent
	if node == t.root {
		t.root = tmp
	} else if node == node.parent.right {
		node.parent.right = tmp
	} else {
		node.parent.left = tmp
	}
	tmp.right = node
	node.parent = tmp
}

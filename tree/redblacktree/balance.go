package redblacktree

// Новый добавляемый элемент всегда является красным
// При добавлений красного баланс может быть нарушен 5 раз
// 1) Добавление корня красного цвета: Просто в конце балансировки всегда красим корень в черный
// 2) Если родитель был красным и (Не надо)дядя красный. Дед не корень: Красим дядю и родителя в черный, деда в красный. Дальше балансируем деда так как родитель деда может оказаться красным
// 3) Если родитель был красным и (Не надо)дядя красный. Дед корень: Делаем то же самое что и 2. Корень в конце балансировки всегда красится в черный
// 4) Если родитель был красным и (Не надо)дядя черный. При этом новый элемент добавляется зигзагом(родитель слева от дяди, новый будет справа от родителя и наобарот): Делаем левый или правый поворот относительно родителя и переносим состояние в 5й случай
// 5) Если родитель был красным и (Не надо)дядя черный. При этом новый элемент добавляется прямолинейно(родитель слева от дяди, новый будет слева от родителя и наобарот): Красим родителя в черный. Деда делаем красным и делаем правый или левый поворот относительно него

func (t *RedBlackTree[T, K]) balanceInsert(newNode *Node[T, K]) {
	var uncle *Node[T, K]
	for newNode.parent.IsRed() { // Убеждаемся что родитель красный
		if newNode.parent.isLeftNode() { // Убеждаемся, что родитель находится с левой стороны от деда
			uncle = newNode.parent.parent.right // Находим дядю
			// Пункты 2 и 3
			if uncle.IsRed() {
				newNode.parent.color = black
				uncle.color = black
				newNode.parent.parent.color = red
				newNode = newNode.parent.parent
			} else { // Пункт 4 и 5
				// Пункт 4
				if newNode.isRightNode() {
					newNode = newNode.parent
					t.leftRotate(newNode)
				}
				// Пункт 5
				newNode.parent.color = black
				newNode.parent.parent.color = red
				t.rightRotate(newNode.parent.parent)
			}
		} else { // Родитель справа от деда
			uncle = newNode.parent.parent.left // Находим дядю
			// Пункты 2 и 3
			if uncle.IsRed() {
				newNode.parent.color = black
				uncle.color = black
				newNode.parent.parent.color = red
				newNode = newNode.parent.parent
			} else { // Пункт 4 и 5
				// Пункт 4
				if newNode.isLeftNode() {
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

func (t *RedBlackTree[T, K]) leftRotate(node *Node[T, K]) {
	tmp := node.right
	node.right = tmp.left
	if node.right.NotNilNode() {
		node.right.parent = node
	}
	tmp.parent = node.parent
	switch node {
	case t.root:
		t.root = tmp
	case node.parent.left:
		node.parent.left = tmp
	default:
		node.parent.right = tmp
	}
	tmp.left = node
	node.parent = tmp
}

func (t *RedBlackTree[T, K]) rightRotate(node *Node[T, K]) {
	tmp := node.left
	node.left = tmp.right
	if node.left.NotNilNode() {
		node.left.parent = node
	}
	tmp.parent = node.parent
	switch node {
	case t.root:
		t.root = tmp
	case node.parent.right:
		node.parent.right = tmp
	default:
		node.parent.left = tmp
	}
	tmp.right = node
	node.parent = tmp
}

func (t *RedBlackTree[T, K]) swapNode(to, from *Node[T, K]) {
	switch to {
	case t.root:
		t.root = from
	case to.parent.left:
		to.parent.left = from
	default:
		to.parent.right = from
	}
	from.parent = to.parent
}

// I) Если брат удаляемого узла был черным:
// 1) брат с двумя узлами красного цвета: Брата красим в цвет родителя, родителя и красного ребенка(для брата справа от родителя это правый для левого наобарот) в черный. Делаем левый или правый поворот относительно родителя
// 2) брат с двумя узлами черного цвета: Это пораждает еще 3 случая:
// 2-1) Если родитель черный корень: Красим брата в красный
// 2-2) Если родитель красный не корень: Красим брата в красный а родителя в черный
// 2-3) Если родитель черный не корень: Родителя красим в красный левый или правый поворот относительно родителя
// 3) брат с черным  слева и красным справа: Решается так же как и первый случай
// 4) брат с черным  справа и красным слева: Брата красим в красный цвет а левого ребенка в черный правый поворот относительно брата. Таким образом мы приходим 1 или 3 случай и решаем так же
// II) Если брат удаляемого узла был красным: Брата в черный родителя в красный левый или правый поворот относительно родителя

func (t *RedBlackTree[T, K]) balanceRemove(node *Node[T, K]) {
	for node != t.root && node.IsBlack() {
		var brother *Node[T, K]
		if node.isLeftNode() { // Если узел слева от родителя
			brother = node.parent.right
			if brother.IsRed() {
				brother.color = black
				node.parent.color = red
				t.leftRotate(node.parent)
				brother = node.parent.right
			}
			if brother.left.IsBlack() && brother.right.IsBlack() { // Случай I-2
				brother.color = red
				node = node.parent // Если родитель корень или красный цикл прервется и он покрасится в черный или будет смотреть другие случаи для родителя
			} else { // Случай I-1 I-3 I-4
				if brother.right.IsBlack() { // Случай I-4
					brother.left.color = black
					brother.color = red
					t.rightRotate(brother)
					brother = node.parent.right
				}
				// Случаи I-1 и I-3
				brother.color = brother.parent.color
				brother.right.color = black
				node.parent.color = black
				t.leftRotate(brother.parent)
				node = t.root
			}
		} else {
			brother = node.parent.left
			if brother.IsRed() {
				brother.color = black
				node.parent.color = red
				t.rightRotate(node.parent)
				brother = node.parent.left
			}
			if brother.left.IsBlack() && brother.right.IsBlack() { // Случай I-2
				brother.color = red
				node = node.parent // Если родитель корень или красный цикл прервется и он покрасится в черный или будет смотреть другие случаи для родителя
			} else { // Случай I-1 I-3 I-4
				if brother.left.IsBlack() { // Случай I-4
					brother.right.color = black
					brother.color = red
					t.leftRotate(brother)
					brother = node.parent.left
				}
				// Случаи I-1 и I-3
				brother.color = brother.parent.color
				brother.left.color = black
				node.parent.color = black
				t.rightRotate(node.parent)
				node = t.root
			}
		}
	}
	node.color = black
}

func (t *RedBlackTree[T, K]) getRightSwappedNode(node *Node[T, K]) *Node[T, K] {
	current := node
	for current.NotNilNode() {
		switch current.getChildrenCount() {
		case 0:
			return current
		case 2:
			if t.cmp(current.left.key, current.key) < 0 {
				current = current.left
			} else {
				current = current.right
			}
		default:
			if current.right.NilNode() {
				if t.cmp(current.left.key, current.key) < 0 {
					current = current.left
				} else {
					return current
				}
			} else {
				if t.cmp(current.right.key, current.key) < 0 {
					current = current.left
				} else {
					return current
				}
			}
		}
	}
	return node
}

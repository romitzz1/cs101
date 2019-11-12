package bag

// Bag is a simple abstract data type (ADT) for storing items. It has methods
// for adding and removing items, checking if a given item exists in the bag,
// and getting the number of items stored in the bag.
type Bag struct {
	// TODO: Add fields here for storing items. You can use any data structure
	// you want (there is no "right" answer). The fields you add here will be
	// accessible in the methods below. See: https://gobyexample.com/methods
	items []int
	amount int
}

// NewBag is used to create an instance of a Bag.
func NewBag() *Bag {
	return &Bag{}
}

// Add inserts an item into the bag. The bag should accept and store duplicate
// items (i.e. it is not a set).
func (b *Bag) Add(i int) {
	b.items = append(b.items, i)
	b.amount++
	return
}

// Remove takes an item out of the bag, if it exists. If there are duplicate
// items in the bag, it should only remove one of them. If the item does not
// exist in the bag, it should silently do nothing.
func (b *Bag) Remove(i int) {
	for index := 0; index < len(b.items); index++ {
		if (b.items[index] == i){
			b.items = append(b.items[:index],b.items[index+1:]...)
			b.amount--
		}
	}
	return
}

// Contains returns true if the given item exists in the bag, and false
// otherwise.
func (b *Bag) Contains(i int) bool {
  for index := 0; index < len(b.items); index++ {
	  if (b.items[index] == i){
		  return true
	  }
  }

	return false
}

// Size returns the number of items in the bag.
func (b *Bag) Size() int {
	return b.amount
}

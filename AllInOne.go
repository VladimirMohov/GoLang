package main

import (
    "math"
    "errors"
    "fmt"
)

const Pi float32 = 3.1415

type position struct {
    x int
    y int
}

type employee struct {
    id int
    pos position
    name string
    sex string
    age int
    salary int
};

type memoryStorage struct {
    data map[int]employee;
}

func newMemoryStorage() *memoryStorage {
    return &memoryStorage{
        data: make(map[int]employee),
    }
}

type storage interface {
    insert(e employee) error;
    get(id int) (employee, error);
    delete(id int) error;
}

// Methods struct
// func (element *type(struct)) method(element type()) (type(), type(error)) {
//    return value, error
//}
//* in os memory
func (s *memoryStorage) insert(e employee) error {
    s.data[e.id] = e;
    fmt.Printf("Employee with such id %d has been add\n", e.id);
    return nil;
}

func (s *memoryStorage) get(id int) (employee, error) {
    e, exists := s.data[id];

    if(!exists) {
        return e, errors.New("Employee with such id doesn`t exists");
    }

    return e, nil;
}

func (s *memoryStorage) delete(id int) error {
    delete(s.data, id);

    return nil;
}

func spawnEmployees(s storage) {
    for i := 0; i <= 10; i++ {
        s.insert(employee{id: i, pos: newPosition(0,0), name: "Vladimir", sex: "male", age: 18, salary: 450_000});
    }
}

func step() func() int {
    count := 0;

    return func() int {
        count++;
        return count;
    }
}

func newPosition(x int, y int) position {
    return position{
        x: x,
        y: y,
    }    
}

func newEmployee(id int, pos position, name string, sex string, age int, salary int) employee {
    return employee{
        id: id,
        pos: pos,
        name: name,
        sex: sex,
        age: age,
        salary: salary,
    }
}

func (e employee) getEmployee() string {
    return fmt.Sprintf("Id: %d, Name: %s, sex: %s, age: %d, salary: %d, positionX: %d, positionY: %d\n", e.id ,e.name, e.sex, e.age, e.salary, e.pos.x, e.pos.y);
}

func (e *employee) goLeft(step int) {
    e.pos.x -= step;
}
func (e *employee) goRight(step int) {
    e.pos.x += step;
}
func (e *employee) goUp(step int) {
    e.pos.y += step;
}
func (e *employee) goDown(step int) {
    e.pos.y -= step;
}

func main() {
    //interface (storage)
    mS := newMemoryStorage();
    spawnEmployees(mS);
    empl, err := mS.get(3);
    if err == nil {
        fmt.Printf("%s\n", empl.getEmployee());
    }

    // func() func() { return func(){} }
    // saved local value
    stepPeople := step();
    fmt.Println(stepPeople());
    fmt.Println(stepPeople());
    fmt.Println(stepPeople());

    // struct (e employee)
    employee1 := newEmployee(0, newPosition(0,0), "Vladimir", "male", 18, 450_000);
    employee2 := newEmployee(0, newPosition(0,0), "Danya", "male", 21, 30_000);
    employee3 := newEmployee(0, newPosition(0,0), "Anastasya", "female", 15, 5_500_000);

    // make (type, min, max)
    numbers2 := make([]int, 5,7);

    // map[type(key)]type(value)
    peoples := make(map[string]int);
    peoples["Vladimir"] = 1;
    peoples["Pavel"] = 2;
    peoples["Sasha"] = 3;
    peoples["Igor"] = 4;
    peoples["Denis"] = 5;

    // Slice - make([]type{}) or []type{} 
    // slice with out make()
    // please use make sure
    var numbers = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20};
    var fruits = []string{
        "apple",
        "pineapple",
        "melon",
        "watermelon",
    }

    // use methods struct
    employee1.goRight(5);
    employee2.goRight(10);
    employee3.goRight(15);
    fmt.Printf("%s\n", employee1.getEmployee());
    fmt.Printf("%s\n", employee2.getEmployee());
    fmt.Printf("%s\n", employee3.getEmployee());

    // for loop
    for people, num := range peoples {
        fmt.Printf("People: %s, Number: %d\n", people, num);
    }

    // map find
    // item type(), exists type(boolean) := mapName[key];
    item, exists := peoples["Danya"];

    if !exists {
        fmt.Println("Is not exists");
    } else {
        fmt.Println("Item is: ", item);
    }

    // custom func()
    customPow(5);

    // array/slice methods 
    a := []int{1,2,3,4,5,6,7,8,9,10};

    // array method append(arrayA, arrayB...) concat array
    a = append(a[:5], a[6:]...); // [1, ... ,5, 7, ..., 10]; a[1:3] -> [1, 2]

    fmt.Println(a);

    // len(a interface{}...) - lenght, cap(a interface{}...) - capacity
    fmt.Printf("len: %d, cap: %d\n", len(numbers2), cap(numbers2)); // len -> 5, cap -> 7

    reverceNumbers := reverce(numbers);

    fmt.Println("Reverse array is : ", reverceNumbers);

    fmt.Printf("Len: %d, cap: %d\n", len(fruits), cap(fruits));
    for index, item := range fruits {
        if (index == 2) {
            continue;
        }
        fmt.Printf("index: %d, name: %s\n", index, item);
    }

    fruits = append(fruits, "lemon", "avocado", "banana");

    fmt.Printf("\nLen: %d, cap: %d\n", len(fruits), cap(fruits));
    for index, item := range fruits {
        if (index == 2) {
            continue;
        }
        fmt.Printf("index: %d, name: %s\n", index, item);
    }

    // &elem get memory elem
    x := 2
    p := &x

    fmt.Println("X: ", x)
    fmt.Println("P: ", p)
    fmt.Println("*P: ", *p)

    // * - value memory; !!! - not link on memory
    *p = 5

    fmt.Println("X: ", x)
    fmt.Println("*P: ", *p)

    circleRadius(5)
    circleRadius(0)
}

func circleRadius(radius int) {
    circleArea, err := calcCircleArea(radius)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Printf("Radius circle is : %d\n", radius)
    fmt.Printf("Diametr circle is : %d\n", (radius * radius))
    fmt.Printf("Area circle is : %f32\n", circleArea)
}

func calcCircleArea(radius int) (float32, error) {
    if radius <= 0 {
        return float32(0), errors.New("The radius shouldn`t be negative")
    }
    return float32(radius) * float32(radius) * math.Pi, nil
}

func reverce(array []int) []int {
    size := len(array);
    returnArray := []int{};
    for i := size; i > 0; i-- {
        returnArray = append(returnArray, array[i - 1]);
    }

    return returnArray;
}

func customPow(n int) {
    for i := 1; i <= n; i++ {
        fmt.Printf("2^%d = %g\n", i, float64(math.Pow(float64(2), float64(i))));
    }
}
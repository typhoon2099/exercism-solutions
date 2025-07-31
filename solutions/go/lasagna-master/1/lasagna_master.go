package lasagna

import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePrepTime int) int {
    if averagePrepTime == 0 {
        return len(layers) * 2
    }

    return len(layers) * averagePrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0

    for _, layer := range layers {
        if layer == "noodles" {
            noodles += 50
        }

        if layer == "sauce" {
            sauce += 0.2
        }

    }
    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
    fmt.Println(myList[len(myList) - 1])
    fmt.Println(friendsList[len(friendsList) - 1])
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, factor int) []float64 { 
    scaled := []float64{}

    for _, quantity := range quantities {
        scaled = append(scaled, quantity * float64(factor) / 2.0)
    }

    return scaled
}
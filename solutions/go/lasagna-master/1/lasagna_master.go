package lasagna

func PreparationTime(layers []string, avgTime int) int {
    if avgTime == 0 {
    	avgTime = 2
    }
    return len(layers) * avgTime
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
    sauce := 0.0
    
    for i := 0; i < len(layers) ; i ++ {
        if layers[i] == "sauce" {
    		sauce += 0.2        
        } else if layers[i] == "noodles" {
            noodles += 50
        }
    }

    return noodles, sauce
}

func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList) - 1] = friendList[len(friendList) - 1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	result := make([]float64, len(quantities))
    for i, v := range quantities {
        result[i] = v * float64(portions) / 2.0
    }

    return result
}

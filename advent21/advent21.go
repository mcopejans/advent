package main

import("fmt")

type Player struct {
	hits, damage, armor int
}

func (this *Player) GetsHit(damage int) {
	hit := damage - this.armor
	if (hit <= 0) {
		this.hits--
	} else {
		this.hits-=hit
	}
}


type StoreItem struct {
	kind string
    props []int // Cost damage armor
}

func Play(player1 Player, player2 Player) bool {
	for {
		player2.GetsHit(player1.damage)

		if (player2.hits <= 0) {
			return true
		}

		player1.GetsHit(player2.damage)

		if (player1.hits <= 0) {
			return false
		}
	}
}

func main() {
	var boss Player
	boss.hits = 100
	boss.damage = 8
	boss.armor = 2

    var store []StoreItem
	var item StoreItem
	item.kind = "weapon"
	item.props = []int{8, 4, 0}
	store = append(store, item)
	item.props = []int{10, 5, 0}
	store = append(store, item)
	item.props = []int{25, 6, 0}
	store = append(store, item)
	item.props = []int{40, 7, 0}
	store = append(store, item)
	item.props = []int{74, 8, 0}
	store = append(store, item)

	var astore []StoreItem
	item.kind = "armor"
	item.props = []int{13, 0, 1}
	astore = append(astore, item)
	item.props = []int{31, 0, 2}
	astore = append(astore, item)
	item.props = []int{53, 0, 3}
	astore = append(astore, item)
	item.props = []int{75, 0, 4}
	astore = append(astore, item)
	item.props = []int{102, 0, 5}
	astore = append(astore, item)
	// no armor
	item.props = []int{0, 0, 0}
	astore = append(astore, item)

	var rstore []StoreItem
	item.kind = "rings"
	item.props = []int{25, 1, 0}
	rstore = append(rstore, item)
	item.props = []int{50, 2, 0}
	rstore = append(rstore, item)
	item.props = []int{100, 3, 0}
	rstore = append(rstore, item)
	item.props = []int{20, 0, 1}
	rstore = append(rstore, item)
	item.props = []int{40, 0, 2}
	rstore = append(rstore, item)
	item.props = []int{80, 0, 3}
	rstore = append(rstore, item)
	// no ring
	item.props = []int{0, 0, 0}
	rstore = append(rstore, item)

	var you Player
	you.hits = 100
	you.damage = 
	you.armor
	var min_cost = 99999999999
	var max_cost = 0
	var min_i, min_j, min_k, min_l int
	var max_i, max_j, max_k, max_l int
	for i := range store {
		for j := range astore {
			for k := range rstore {
				for l := range rstore {
					if k == l {
						continue
					}
					you.damage = store[i].props[1] + astore[j].props[1] + rstore[k].props[1] + rstore[l].props[1]
					you.armor  = store[i].props[2] + astore[j].props[2] + rstore[k].props[2] + rstore[l].props[2]

					total_cost := store[i].props[0] + astore[j].props[0] + rstore[k].props[0] + rstore[l].props[0]
					if Play(you, boss) {
						if total_cost < min_cost {
							min_cost = total_cost
							min_i = i
							min_j = j
							min_k = k
							min_l = l
						}
					} else {
						if total_cost > max_cost {
							max_cost = total_cost
							max_i = i
							max_j = j
							max_k = k
							max_l = l
						}
					}
				}
			}
		}
	}

	fmt.Println("Min cost while still winning: ", min_cost)
	fmt.Println("(",min_i,", ",min_j,", ",min_k,", ",min_l,")")
	fmt.Println("Min cost while losing: ", max_cost)
	fmt.Println("(",max_i,", ",max_j,", ",max_k,", ",max_l,")")

	boss.hits=12
	you.hits=8
	boss.damage=7
	you.damage=5
	boss.armor=2
	you.armor=5
	Play(you,boss)
}
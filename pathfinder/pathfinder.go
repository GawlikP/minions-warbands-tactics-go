package pathfinder

import (
  "minions-warbands-tactics/constant"
  "fmt"
  "slices"
  "math"
  "log"
  "errors"
)

type Node struct {
  X       int
  Y       int
  H       int
  G       int
  F       int
	Parent  *Node
  Block   bool
}

func AStar(sx, sy, dx, dy, width int, grid []constant.BattleMapTileType) ([]int, error) {
  Start := Node{
    X: sx,
    Y: sy,
    H: 0,
    G: 0,
    F: 0,
    Parent: nil,
  }
  Dest := Node{
    X: dx,
    Y: dy,
    H: 0,
    G: 0,
    F: 0,
    Parent: nil,
  }
  Nodes := GenerateNodes(grid, width)
  OpenList := []int{}
  ClosedList := []int{}
  PathGenerated := false

  OpenList = append(OpenList, Start.X + Start.Y * width)
  for len(OpenList) > 0 {

    for i, v := range  OpenList {
      log.Printf("OpenList[%d]: X:%d Y:%d", i, v%width, v/width)
    }
    currentNodeIndex := getFMin(OpenList, Nodes)
    if currentNodeIndex == -1 {
      return nil, errors.New("Cannot find path") 
    }
    currentNode := Nodes[currentNodeIndex]
    //remove from open list

    indexToRemove := 0
    for i := 0; i < len(OpenList); i++ {
      if currentNodeIndex == OpenList[i] {
        indexToRemove = i
      }
    }
    OpenList = append(OpenList[:indexToRemove], OpenList[indexToRemove+1:]...)

    //add to close list
    if Nodes[currentNodeIndex].X == dx && Nodes[currentNodeIndex].Y == dy {
      log.Print("Path Found!")
      return generatePath(Nodes[currentNodeIndex], width), nil
    }
    if PathGenerated {
      log.Print("Closing Loop")
      break
    }
    ClosedList = append(ClosedList, currentNodeIndex)

    // get Neighbors
    availableNeighbors := getNeighbors(currentNode, currentNodeIndex, grid, width, Nodes)
    for _, p := range availableNeighbors {
      // UPDATE WEIGHT
      parent := p.Parent
      hor := (int)(math.Abs((float64)(p.X - parent.X)))
      ver := (int)(math.Abs((float64)(p.Y - parent.Y)))
      p.G = hor + ver

      absx := (int)(math.Abs((float64)(Dest.X - p.X)))
      absy := (int)(math.Abs((float64)(Dest.Y - p.Y)))
      p.H = (absx + absy) * 10
      p.F = p.G + p.H
      // UPDATE WEIGHT
      if !slices.Contains(ClosedList, p.X + p.Y * width) {
        if !slices.Contains(OpenList, p.X + p.Y * width) {
          OpenList = append(OpenList, p.X + p.Y * width)
          Nodes[p.X + p.Y * width] = p
        } else {
          if p.F > Nodes[p.X + p.Y * width].F {
            Nodes[p.X + p.Y * width].Parent = &Nodes[currentNodeIndex]
          }
        }
      }

    }
    // grid.Draw()
  }
  return nil, nil 
}


func generatePath(node Node, width int) []int {
  nodesPath := []int{}

  for node.Parent != nil {
    nodesPath = append(nodesPath, node.X + node.Y * width)
    node = *(node.Parent)
  }

  return nodesPath
}


func getNeighbors(node Node, currentNodeIndex int, grid []constant.BattleMapTileType, width int, Nodes []Node) []Node {
	var neighbors []Node
  length := len(grid)
  left := leftIndex(node, width, length)
  right := rightIndex(node, width, length)
  up := upIndex(node, width, length)
  down := downIndex(node, width, length)

  if left != -1 && grid[left] != constant.Stone {
    node := Nodes[left]
    node.Parent = &Nodes[currentNodeIndex]
    neighbors = append(neighbors, node)
  }
  if right != -1 && grid[right] != constant.Stone {
    node := Nodes[right]
    node.Parent = &Nodes[currentNodeIndex]
    neighbors = append(neighbors, node)
  }
  if up != -1 && grid[up] != constant.Stone {
    node := Nodes[up]
    node.Parent = &Nodes[currentNodeIndex]
    neighbors = append(neighbors, node)
  }
  if down != -1 && grid[down] != constant.Stone {
    node := Nodes[down]
    node.Parent = &Nodes[currentNodeIndex]
    neighbors = append(neighbors, node)
  }
	return neighbors
}

func leftIndex(node Node, width int, length int) int {
  left := node.X - 1 + node.Y * width
  if node.X -1 > -1 && left > -1 && left < length {
    return left 
  }
  return -1
}

func rightIndex(node Node, width int, length int) int {
  right := node.X + 1 + node.Y * width
  if node.X + 1 < width && right > -1 && right < length {
    return right
  }
  return -1
}

func upIndex(node Node, width int, length int) int {
  up := node.X + node.Y * width - width
  if node.Y - 1 > -1 && up > -1 &&  up < length {
    return up 
  }
  return -1
}

func downIndex(node Node, width int, length int) int {
  down := node.X + node.Y * width + width
  if node.Y + 1 < length/width && down > -1 &&  down < length {
    return down 
  }
  return -1
}

func GenerateNodes(grid []constant.BattleMapTileType, width int) []Node {
  var nodes []Node
  for i := 0; i < len(grid); i++ {
    node := Node{
      X: i%width,
      Y: i/width,
      H: 0,
      G: 0,
      F: math.MaxInt64,
      Block: false,
      Parent: nil,
    }
    if grid[i] == constant.Stone {
      node.Block = true
    }
    nodes = append(nodes, node)
  }
  return nodes
}

func getFMin(openList []int, nodes []Node) int {
	if len(openList) == 0 {
		fmt.Print("No way!!!")
    return -1
	}
	index := 0
	for i, p := range openList {
    node := nodes[openList[index]]
		if (i > 0) && (nodes[p].F <= node.F) {
			index = i
		}
	}
	return openList[index]
}

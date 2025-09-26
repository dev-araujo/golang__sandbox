package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dev-araujo/golang__sandbox/projects/basic-blockchain-go/basic-blockchain-with-auditor/blockchain"
)

const EXIT_COMMAND = "/exit"
const NEW_DATA_SEPARATOR = "------------------------------"
const BLOCK_SEPARATOR = "=============================="

func main() {
	bc := blockchain.NewBlockchain()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("enter the data for a new block: ")
		data, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("error %v\n", err)
			break
		}

		trimmedData := strings.TrimSpace(data)
		if trimmedData == EXIT_COMMAND {
			break
		}
		if len(trimmedData) > 0 {
			bc.AddBlock([]byte(trimmedData))
			if !bc.Auditor() {
				fmt.Printf("REDE COMPROMETIDA NO BLOCO:\n %v  ", bc.Blocks[len(bc.Blocks)-1])
			}
			fmt.Println(NEW_DATA_SEPARATOR)
			printInline(bc)
		}

	}

}

func printInline(bc *blockchain.Blockchain) {
	for _, block := range bc.Blocks {
		fmt.Println(BLOCK_SEPARATOR)
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println(BLOCK_SEPARATOR)
	}
}

package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

//Напишите построчную сортировку большого текстового файла, не влезающего в память.
//Размер требуемой памяти не должен зависеть от размера файла.
//Длина строки разумная, одна строка сильно меньше, чем объем памяти.
//Для проверки работоспособности, нужен генератор таких файлов, принимающий
//в качестве параметров количество строки их максимальную длину.
//Язык программирования не важен. На задание рекомендуется потратить не более 4 часов.

const (
	chunkSize = 10
)

func main() {
	sourceFile, err := os.Open("./joom/testFile.txt")
	if err != nil {
		log.Fatalf("failed opening source file: %s", err)
	}
	defer sourceFile.Close()

	scanner := bufio.NewScanner(sourceFile)
	scanner.Split(bufio.ScanLines)
	var lines []string

	chunkNum := 0
	for {
		for i := 0; i < chunkSize; i++ {
			if !scanner.Scan() {
				break
			}
			lines = append(lines, scanner.Text())
		}

		if len(lines) == 0 {
			break
		}

		sortChunk(lines, chunkNum)

		lines = lines[:0]
		chunkNum++
	}

	mergeChunks(chunkNum)

	return
}

func mergeChunks(chunkNum int) {

	scanners := make([]*bufio.Scanner, 0, chunkNum)

	for i := 0; i < chunkNum; i++ {
		chunkFile, err := os.Open(getFileNameForChunk(i))
		if err != nil {
			log.Fatalf("failed openning chunk file: %s", err)
		}
		defer chunkFile.Close()

		scanner := bufio.NewScanner(chunkFile)
		scanners = append(scanners, scanner)
	}

	resFile, err := os.Create("./joom/sorted.txt")
	if err != nil {
		log.Fatalf("failed openning result file: %s", err)
	}

	type line struct {
		value string
		done  bool
	}
	lines := make([]line, 0, chunkNum)
	for _, sc := range scanners {
		done := !sc.Scan()
		value := sc.Text()
		lines = append(lines, line{
			value: value,
			done:  done,
		})
	}

	for {
		minLineInd := -1
		i := 0
		for ; i < len(lines); i++ {
			if !lines[i].done {
				minLineInd = i
				break
			}
		}
		if minLineInd == -1 {
			return
		}

		for ; i < len(lines); i++ {
			if !lines[i].done && lines[i].value < lines[minLineInd].value {
				minLineInd = i
			}
		}

		if _, err := resFile.Write([]byte(lines[minLineInd].value)); err != nil && err != io.EOF {
			log.Fatalf("failed writing to file: %s", err)
		}
		if _, err := resFile.Write([]byte("\n")); err != nil && err != io.EOF {
			log.Fatalf("failed writing separator to file: %s", err)
		}
		lines[minLineInd].done = !scanners[minLineInd].Scan()
		lines[minLineInd].value = scanners[minLineInd].Text()
	}
}

func sortChunk(lines []string, chunkNum int) {

	sort.Strings(lines)

	destFile, err := os.Create(getFileNameForChunk(chunkNum))
	if err != nil {
		log.Fatalf("failed opening destination file: %s", err)
	}

	for i := range lines {
		if _, err := destFile.Write([]byte(lines[i])); err != nil && err != io.EOF {
			log.Fatalf("failed writing to file: %s", err)
		}
		if _, err := destFile.Write([]byte("\n")); err != nil && err != io.EOF {
			log.Fatalf("failed writing separator to file: %s", err)
		}
	}
}

func getFileNameForChunk(n int) string {
	return "./joom/chunk_" + strconv.Itoa(n) + ".txt"
}

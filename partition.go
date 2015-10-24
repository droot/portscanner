package scanner

func partition(slice []uint16, numParts int) [][]uint16 {

  var parts [][]uint16

  if numParts <= 0 || len(slice) == 0 {
    return nil
  }

  partSize := len(slice) / numParts

  var ii int
  for ii = 0; ii < numParts; ii++ {
    parts = append(parts, slice[ii*partSize:(ii+1)*partSize])
  }

  if len(slice) > (numParts * partSize) {
    lastPart := parts[numParts-1]
    lastPart = append(lastPart, slice[numParts*partSize:]...)
    parts[numParts-1] = lastPart
  }

  return parts
}
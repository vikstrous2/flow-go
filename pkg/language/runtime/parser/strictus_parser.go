// Code generated from parser/Strictus.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Strictus
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "strings"

var _ = strings.Builder{}

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 72, 768,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 70, 4,
	71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75, 4, 76,
	9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4, 81, 9,
	81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86, 9, 86,
	4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 3, 2, 3, 2, 5, 2, 181, 10, 2,
	7, 2, 183, 10, 2, 12, 2, 14, 2, 186, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5,
	3, 192, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 200, 10, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 7, 5, 206, 10, 5, 12, 5, 14, 5, 209, 11, 5, 3, 5,
	3, 5, 5, 5, 213, 10, 5, 3, 5, 3, 5, 5, 5, 217, 10, 5, 3, 6, 3, 6, 3, 6,
	5, 6, 222, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 7, 8, 236, 10, 8, 12, 8, 14, 8, 239, 11, 8, 5, 8, 241,
	10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 247, 10, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 7,
	12, 262, 10, 12, 12, 12, 14, 12, 265, 11, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 5, 13, 272, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 5, 15, 279,
	10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16,
	289, 10, 16, 3, 16, 5, 16, 292, 10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 304, 10, 18, 12, 18, 14,
	18, 307, 11, 18, 5, 18, 309, 10, 18, 3, 18, 3, 18, 3, 19, 5, 19, 314, 10,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 5, 20, 321, 10, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 331, 10, 21, 12, 21, 14,
	21, 334, 11, 21, 5, 21, 336, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	5, 22, 343, 10, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7,
	24, 352, 10, 24, 12, 24, 14, 24, 355, 11, 24, 5, 24, 357, 10, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 5, 29, 386, 10, 29, 3, 29, 5, 29, 389,
	10, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 7, 32, 407, 10, 32, 12,
	32, 14, 32, 410, 11, 32, 3, 33, 3, 33, 3, 33, 5, 33, 415, 10, 33, 3, 34,
	3, 34, 3, 34, 7, 34, 420, 10, 34, 12, 34, 14, 34, 423, 11, 34, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 435,
	10, 35, 3, 36, 3, 36, 3, 36, 5, 36, 440, 10, 36, 3, 37, 3, 37, 3, 38, 3,
	38, 3, 39, 3, 39, 3, 39, 5, 39, 449, 10, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	5, 39, 455, 10, 39, 5, 39, 457, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 471, 10, 42,
	3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 7, 43, 478, 10, 43, 12, 43, 14, 43,
	481, 11, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3,
	45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 5, 47, 500,
	10, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 7, 48, 508, 10, 48, 12,
	48, 14, 48, 511, 11, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 7, 49,
	519, 10, 49, 12, 49, 14, 49, 522, 11, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3,
	50, 3, 50, 3, 50, 7, 50, 531, 10, 50, 12, 50, 14, 50, 534, 11, 50, 3, 51,
	3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 7, 51, 543, 10, 51, 12, 51, 14,
	51, 546, 11, 51, 3, 52, 3, 52, 3, 52, 5, 52, 551, 10, 52, 3, 53, 3, 53,
	3, 53, 3, 53, 3, 53, 3, 53, 7, 53, 559, 10, 53, 12, 53, 14, 53, 562, 11,
	53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 7, 54, 570, 10, 54, 12, 54,
	14, 54, 573, 11, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 7,
	55, 582, 10, 55, 12, 55, 14, 55, 585, 11, 55, 3, 56, 3, 56, 3, 56, 3, 56,
	3, 56, 3, 56, 3, 56, 7, 56, 594, 10, 56, 12, 56, 14, 56, 597, 11, 56, 3,
	57, 3, 57, 6, 57, 601, 10, 57, 13, 57, 14, 57, 602, 3, 57, 3, 57, 5, 57,
	607, 10, 57, 3, 58, 3, 58, 3, 58, 3, 58, 5, 58, 613, 10, 58, 3, 59, 3,
	59, 7, 59, 617, 10, 59, 12, 59, 14, 59, 620, 11, 59, 3, 60, 3, 60, 5, 60,
	624, 10, 60, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 3,
	65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 5, 66, 640, 10, 66, 3, 67, 3, 67,
	3, 67, 3, 67, 3, 68, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3,
	70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 72, 5, 72, 662, 10, 72,
	3, 72, 3, 72, 3, 73, 3, 73, 3, 73, 3, 73, 3, 74, 3, 74, 5, 74, 672, 10,
	74, 3, 75, 3, 75, 3, 75, 3, 76, 3, 76, 3, 76, 5, 76, 680, 10, 76, 3, 76,
	3, 76, 3, 77, 3, 77, 3, 77, 3, 77, 7, 77, 688, 10, 77, 12, 77, 14, 77,
	691, 11, 77, 5, 77, 693, 10, 77, 3, 77, 3, 77, 3, 78, 3, 78, 3, 78, 5,
	78, 700, 10, 78, 3, 78, 3, 78, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79,
	5, 79, 710, 10, 79, 3, 80, 3, 80, 3, 81, 3, 81, 3, 82, 3, 82, 3, 83, 5,
	83, 719, 10, 83, 3, 83, 3, 83, 3, 84, 3, 84, 3, 84, 3, 84, 3, 84, 5, 84,
	728, 10, 84, 3, 85, 3, 85, 3, 85, 3, 85, 7, 85, 734, 10, 85, 12, 85, 14,
	85, 737, 11, 85, 5, 85, 739, 10, 85, 3, 85, 3, 85, 3, 86, 3, 86, 3, 86,
	3, 86, 7, 86, 747, 10, 86, 12, 86, 14, 86, 750, 11, 86, 5, 86, 752, 10,
	86, 3, 86, 3, 86, 3, 87, 3, 87, 3, 87, 3, 87, 3, 88, 3, 88, 3, 89, 3, 89,
	3, 89, 3, 89, 5, 89, 766, 10, 89, 3, 89, 2, 10, 94, 96, 98, 100, 104, 106,
	108, 110, 90, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
	70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
	106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134,
	136, 138, 140, 142, 144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164,
	166, 168, 170, 172, 174, 176, 2, 12, 3, 2, 50, 51, 3, 2, 36, 38, 4, 2,
	11, 11, 28, 28, 3, 2, 15, 16, 3, 2, 17, 20, 3, 2, 21, 22, 3, 2, 23, 25,
	4, 2, 22, 22, 27, 28, 3, 2, 55, 56, 3, 2, 59, 62, 2, 778, 2, 184, 3, 2,
	2, 2, 4, 191, 3, 2, 2, 2, 6, 199, 3, 2, 2, 2, 8, 201, 3, 2, 2, 2, 10, 221,
	3, 2, 2, 2, 12, 223, 3, 2, 2, 2, 14, 240, 3, 2, 2, 2, 16, 242, 3, 2, 2,
	2, 18, 244, 3, 2, 2, 2, 20, 252, 3, 2, 2, 2, 22, 263, 3, 2, 2, 2, 24, 271,
	3, 2, 2, 2, 26, 273, 3, 2, 2, 2, 28, 275, 3, 2, 2, 2, 30, 282, 3, 2, 2,
	2, 32, 295, 3, 2, 2, 2, 34, 299, 3, 2, 2, 2, 36, 313, 3, 2, 2, 2, 38, 320,
	3, 2, 2, 2, 40, 335, 3, 2, 2, 2, 42, 342, 3, 2, 2, 2, 44, 344, 3, 2, 2,
	2, 46, 346, 3, 2, 2, 2, 48, 363, 3, 2, 2, 2, 50, 367, 3, 2, 2, 2, 52, 373,
	3, 2, 2, 2, 54, 379, 3, 2, 2, 2, 56, 383, 3, 2, 2, 2, 58, 393, 3, 2, 2,
	2, 60, 398, 3, 2, 2, 2, 62, 408, 3, 2, 2, 2, 64, 411, 3, 2, 2, 2, 66, 421,
	3, 2, 2, 2, 68, 434, 3, 2, 2, 2, 70, 436, 3, 2, 2, 2, 72, 441, 3, 2, 2,
	2, 74, 443, 3, 2, 2, 2, 76, 445, 3, 2, 2, 2, 78, 458, 3, 2, 2, 2, 80, 462,
	3, 2, 2, 2, 82, 466, 3, 2, 2, 2, 84, 475, 3, 2, 2, 2, 86, 485, 3, 2, 2,
	2, 88, 489, 3, 2, 2, 2, 90, 491, 3, 2, 2, 2, 92, 493, 3, 2, 2, 2, 94, 501,
	3, 2, 2, 2, 96, 512, 3, 2, 2, 2, 98, 523, 3, 2, 2, 2, 100, 535, 3, 2, 2,
	2, 102, 547, 3, 2, 2, 2, 104, 552, 3, 2, 2, 2, 106, 563, 3, 2, 2, 2, 108,
	574, 3, 2, 2, 2, 110, 586, 3, 2, 2, 2, 112, 606, 3, 2, 2, 2, 114, 612,
	3, 2, 2, 2, 116, 614, 3, 2, 2, 2, 118, 623, 3, 2, 2, 2, 120, 625, 3, 2,
	2, 2, 122, 627, 3, 2, 2, 2, 124, 629, 3, 2, 2, 2, 126, 631, 3, 2, 2, 2,
	128, 633, 3, 2, 2, 2, 130, 639, 3, 2, 2, 2, 132, 641, 3, 2, 2, 2, 134,
	645, 3, 2, 2, 2, 136, 648, 3, 2, 2, 2, 138, 653, 3, 2, 2, 2, 140, 655,
	3, 2, 2, 2, 142, 657, 3, 2, 2, 2, 144, 665, 3, 2, 2, 2, 146, 671, 3, 2,
	2, 2, 148, 673, 3, 2, 2, 2, 150, 676, 3, 2, 2, 2, 152, 683, 3, 2, 2, 2,
	154, 699, 3, 2, 2, 2, 156, 709, 3, 2, 2, 2, 158, 711, 3, 2, 2, 2, 160,
	713, 3, 2, 2, 2, 162, 715, 3, 2, 2, 2, 164, 718, 3, 2, 2, 2, 166, 727,
	3, 2, 2, 2, 168, 729, 3, 2, 2, 2, 170, 742, 3, 2, 2, 2, 172, 755, 3, 2,
	2, 2, 174, 759, 3, 2, 2, 2, 176, 765, 3, 2, 2, 2, 178, 180, 5, 6, 4, 2,
	179, 181, 7, 3, 2, 2, 180, 179, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181,
	183, 3, 2, 2, 2, 182, 178, 3, 2, 2, 2, 183, 186, 3, 2, 2, 2, 184, 182,
	3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 187, 3, 2, 2, 2, 186, 184, 3, 2,
	2, 2, 187, 188, 7, 2, 2, 3, 188, 3, 3, 2, 2, 2, 189, 192, 5, 2, 2, 2, 190,
	192, 5, 66, 34, 2, 191, 189, 3, 2, 2, 2, 191, 190, 3, 2, 2, 2, 192, 5,
	3, 2, 2, 2, 193, 200, 5, 12, 7, 2, 194, 200, 5, 20, 11, 2, 195, 200, 5,
	30, 16, 2, 196, 200, 5, 82, 42, 2, 197, 200, 5, 8, 5, 2, 198, 200, 5, 32,
	17, 2, 199, 193, 3, 2, 2, 2, 199, 194, 3, 2, 2, 2, 199, 195, 3, 2, 2, 2,
	199, 196, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 198, 3, 2, 2, 2, 200,
	7, 3, 2, 2, 2, 201, 212, 7, 58, 2, 2, 202, 207, 5, 174, 88, 2, 203, 204,
	7, 4, 2, 2, 204, 206, 5, 174, 88, 2, 205, 203, 3, 2, 2, 2, 206, 209, 3,
	2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 210, 3, 2, 2,
	2, 209, 207, 3, 2, 2, 2, 210, 211, 7, 59, 2, 2, 211, 213, 3, 2, 2, 2, 212,
	202, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 217,
	5, 162, 82, 2, 215, 217, 7, 66, 2, 2, 216, 214, 3, 2, 2, 2, 216, 215, 3,
	2, 2, 2, 217, 9, 3, 2, 2, 2, 218, 222, 3, 2, 2, 2, 219, 222, 7, 45, 2,
	2, 220, 222, 7, 46, 2, 2, 221, 218, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221,
	220, 3, 2, 2, 2, 222, 11, 3, 2, 2, 2, 223, 224, 5, 10, 6, 2, 224, 225,
	5, 26, 14, 2, 225, 226, 5, 174, 88, 2, 226, 227, 5, 14, 8, 2, 227, 228,
	7, 5, 2, 2, 228, 229, 5, 22, 12, 2, 229, 230, 7, 6, 2, 2, 230, 13, 3, 2,
	2, 2, 231, 232, 7, 7, 2, 2, 232, 237, 5, 174, 88, 2, 233, 234, 7, 4, 2,
	2, 234, 236, 5, 174, 88, 2, 235, 233, 3, 2, 2, 2, 236, 239, 3, 2, 2, 2,
	237, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 241, 3, 2, 2, 2, 239,
	237, 3, 2, 2, 2, 240, 231, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 15, 3,
	2, 2, 2, 242, 243, 9, 2, 2, 2, 243, 17, 3, 2, 2, 2, 244, 246, 5, 10, 6,
	2, 245, 247, 5, 16, 9, 2, 246, 245, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247,
	248, 3, 2, 2, 2, 248, 249, 5, 174, 88, 2, 249, 250, 7, 7, 2, 2, 250, 251,
	5, 38, 20, 2, 251, 19, 3, 2, 2, 2, 252, 253, 5, 10, 6, 2, 253, 254, 5,
	26, 14, 2, 254, 255, 7, 39, 2, 2, 255, 256, 5, 174, 88, 2, 256, 257, 7,
	5, 2, 2, 257, 258, 5, 22, 12, 2, 258, 259, 7, 6, 2, 2, 259, 21, 3, 2, 2,
	2, 260, 262, 5, 24, 13, 2, 261, 260, 3, 2, 2, 2, 262, 265, 3, 2, 2, 2,
	263, 261, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 23, 3, 2, 2, 2, 265, 263,
	3, 2, 2, 2, 266, 272, 5, 18, 10, 2, 267, 272, 5, 28, 15, 2, 268, 272, 5,
	30, 16, 2, 269, 272, 5, 20, 11, 2, 270, 272, 5, 12, 7, 2, 271, 266, 3,
	2, 2, 2, 271, 267, 3, 2, 2, 2, 271, 268, 3, 2, 2, 2, 271, 269, 3, 2, 2,
	2, 271, 270, 3, 2, 2, 2, 272, 25, 3, 2, 2, 2, 273, 274, 9, 3, 2, 2, 274,
	27, 3, 2, 2, 2, 275, 276, 5, 174, 88, 2, 276, 278, 5, 34, 18, 2, 277, 279,
	5, 56, 29, 2, 278, 277, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 280, 3,
	2, 2, 2, 280, 281, 6, 15, 2, 3, 281, 29, 3, 2, 2, 2, 282, 283, 5, 10, 6,
	2, 283, 284, 7, 40, 2, 2, 284, 285, 5, 174, 88, 2, 285, 288, 5, 34, 18,
	2, 286, 287, 7, 7, 2, 2, 287, 289, 5, 38, 20, 2, 288, 286, 3, 2, 2, 2,
	288, 289, 3, 2, 2, 2, 289, 291, 3, 2, 2, 2, 290, 292, 5, 56, 29, 2, 291,
	290, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 294,
	6, 16, 3, 3, 294, 31, 3, 2, 2, 2, 295, 296, 7, 41, 2, 2, 296, 297, 5, 174,
	88, 2, 297, 298, 5, 34, 18, 2, 298, 33, 3, 2, 2, 2, 299, 308, 7, 33, 2,
	2, 300, 305, 5, 36, 19, 2, 301, 302, 7, 4, 2, 2, 302, 304, 5, 36, 19, 2,
	303, 301, 3, 2, 2, 2, 304, 307, 3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 305,
	306, 3, 2, 2, 2, 306, 309, 3, 2, 2, 2, 307, 305, 3, 2, 2, 2, 308, 300,
	3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 311, 7, 34,
	2, 2, 311, 35, 3, 2, 2, 2, 312, 314, 5, 174, 88, 2, 313, 312, 3, 2, 2,
	2, 313, 314, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 316, 5, 174, 88, 2,
	316, 317, 7, 7, 2, 2, 317, 318, 5, 38, 20, 2, 318, 37, 3, 2, 2, 2, 319,
	321, 7, 28, 2, 2, 320, 319, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321, 322,
	3, 2, 2, 2, 322, 323, 5, 40, 21, 2, 323, 39, 3, 2, 2, 2, 324, 325, 7, 26,
	2, 2, 325, 326, 6, 21, 4, 2, 326, 336, 5, 42, 22, 2, 327, 332, 5, 42, 22,
	2, 328, 329, 6, 21, 5, 2, 329, 331, 7, 29, 2, 2, 330, 328, 3, 2, 2, 2,
	331, 334, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333,
	336, 3, 2, 2, 2, 334, 332, 3, 2, 2, 2, 335, 324, 3, 2, 2, 2, 335, 327,
	3, 2, 2, 2, 336, 41, 3, 2, 2, 2, 337, 343, 5, 44, 23, 2, 338, 343, 5, 46,
	24, 2, 339, 343, 5, 48, 25, 2, 340, 343, 5, 50, 26, 2, 341, 343, 5, 52,
	27, 2, 342, 337, 3, 2, 2, 2, 342, 338, 3, 2, 2, 2, 342, 339, 3, 2, 2, 2,
	342, 340, 3, 2, 2, 2, 342, 341, 3, 2, 2, 2, 343, 43, 3, 2, 2, 2, 344, 345,
	5, 174, 88, 2, 345, 45, 3, 2, 2, 2, 346, 347, 7, 33, 2, 2, 347, 356, 7,
	33, 2, 2, 348, 353, 5, 38, 20, 2, 349, 350, 7, 4, 2, 2, 350, 352, 5, 38,
	20, 2, 351, 349, 3, 2, 2, 2, 352, 355, 3, 2, 2, 2, 353, 351, 3, 2, 2, 2,
	353, 354, 3, 2, 2, 2, 354, 357, 3, 2, 2, 2, 355, 353, 3, 2, 2, 2, 356,
	348, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357, 358, 3, 2, 2, 2, 358, 359,
	7, 34, 2, 2, 359, 360, 7, 7, 2, 2, 360, 361, 5, 38, 20, 2, 361, 362, 7,
	34, 2, 2, 362, 47, 3, 2, 2, 2, 363, 364, 7, 8, 2, 2, 364, 365, 5, 40, 21,
	2, 365, 366, 7, 9, 2, 2, 366, 49, 3, 2, 2, 2, 367, 368, 7, 8, 2, 2, 368,
	369, 5, 40, 21, 2, 369, 370, 7, 3, 2, 2, 370, 371, 7, 63, 2, 2, 371, 372,
	7, 9, 2, 2, 372, 51, 3, 2, 2, 2, 373, 374, 7, 5, 2, 2, 374, 375, 5, 40,
	21, 2, 375, 376, 7, 7, 2, 2, 376, 377, 5, 40, 21, 2, 377, 378, 7, 6, 2,
	2, 378, 53, 3, 2, 2, 2, 379, 380, 7, 5, 2, 2, 380, 381, 5, 66, 34, 2, 381,
	382, 7, 6, 2, 2, 382, 55, 3, 2, 2, 2, 383, 385, 7, 5, 2, 2, 384, 386, 5,
	58, 30, 2, 385, 384, 3, 2, 2, 2, 385, 386, 3, 2, 2, 2, 386, 388, 3, 2,
	2, 2, 387, 389, 5, 60, 31, 2, 388, 387, 3, 2, 2, 2, 388, 389, 3, 2, 2,
	2, 389, 390, 3, 2, 2, 2, 390, 391, 5, 66, 34, 2, 391, 392, 7, 6, 2, 2,
	392, 57, 3, 2, 2, 2, 393, 394, 7, 43, 2, 2, 394, 395, 7, 5, 2, 2, 395,
	396, 5, 62, 32, 2, 396, 397, 7, 6, 2, 2, 397, 59, 3, 2, 2, 2, 398, 399,
	7, 44, 2, 2, 399, 400, 7, 5, 2, 2, 400, 401, 5, 62, 32, 2, 401, 402, 7,
	6, 2, 2, 402, 61, 3, 2, 2, 2, 403, 404, 5, 64, 33, 2, 404, 405, 5, 176,
	89, 2, 405, 407, 3, 2, 2, 2, 406, 403, 3, 2, 2, 2, 407, 410, 3, 2, 2, 2,
	408, 406, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 63, 3, 2, 2, 2, 410, 408,
	3, 2, 2, 2, 411, 414, 5, 90, 46, 2, 412, 413, 7, 7, 2, 2, 413, 415, 5,
	90, 46, 2, 414, 412, 3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 65, 3, 2, 2,
	2, 416, 417, 5, 68, 35, 2, 417, 418, 5, 176, 89, 2, 418, 420, 3, 2, 2,
	2, 419, 416, 3, 2, 2, 2, 420, 423, 3, 2, 2, 2, 421, 419, 3, 2, 2, 2, 421,
	422, 3, 2, 2, 2, 422, 67, 3, 2, 2, 2, 423, 421, 3, 2, 2, 2, 424, 435, 5,
	70, 36, 2, 425, 435, 5, 72, 37, 2, 426, 435, 5, 74, 38, 2, 427, 435, 5,
	76, 39, 2, 428, 435, 5, 78, 40, 2, 429, 435, 5, 80, 41, 2, 430, 435, 5,
	6, 4, 2, 431, 435, 5, 84, 43, 2, 432, 435, 5, 86, 44, 2, 433, 435, 5, 90,
	46, 2, 434, 424, 3, 2, 2, 2, 434, 425, 3, 2, 2, 2, 434, 426, 3, 2, 2, 2,
	434, 427, 3, 2, 2, 2, 434, 428, 3, 2, 2, 2, 434, 429, 3, 2, 2, 2, 434,
	430, 3, 2, 2, 2, 434, 431, 3, 2, 2, 2, 434, 432, 3, 2, 2, 2, 434, 433,
	3, 2, 2, 2, 435, 69, 3, 2, 2, 2, 436, 439, 7, 47, 2, 2, 437, 438, 6, 36,
	6, 2, 438, 440, 5, 90, 46, 2, 439, 437, 3, 2, 2, 2, 439, 440, 3, 2, 2,
	2, 440, 71, 3, 2, 2, 2, 441, 442, 7, 48, 2, 2, 442, 73, 3, 2, 2, 2, 443,
	444, 7, 49, 2, 2, 444, 75, 3, 2, 2, 2, 445, 448, 7, 52, 2, 2, 446, 449,
	5, 90, 46, 2, 447, 449, 5, 82, 42, 2, 448, 446, 3, 2, 2, 2, 448, 447, 3,
	2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 456, 5, 54, 28, 2, 451, 454, 7, 53,
	2, 2, 452, 455, 5, 76, 39, 2, 453, 455, 5, 54, 28, 2, 454, 452, 3, 2, 2,
	2, 454, 453, 3, 2, 2, 2, 455, 457, 3, 2, 2, 2, 456, 451, 3, 2, 2, 2, 456,
	457, 3, 2, 2, 2, 457, 77, 3, 2, 2, 2, 458, 459, 7, 54, 2, 2, 459, 460,
	5, 90, 46, 2, 460, 461, 5, 54, 28, 2, 461, 79, 3, 2, 2, 2, 462, 463, 7,
	42, 2, 2, 463, 464, 5, 174, 88, 2, 464, 465, 5, 152, 77, 2, 465, 81, 3,
	2, 2, 2, 466, 467, 5, 16, 9, 2, 467, 470, 5, 174, 88, 2, 468, 469, 7, 7,
	2, 2, 469, 471, 5, 38, 20, 2, 470, 468, 3, 2, 2, 2, 470, 471, 3, 2, 2,
	2, 471, 472, 3, 2, 2, 2, 472, 473, 5, 88, 45, 2, 473, 474, 5, 90, 46, 2,
	474, 83, 3, 2, 2, 2, 475, 479, 5, 174, 88, 2, 476, 478, 5, 146, 74, 2,
	477, 476, 3, 2, 2, 2, 478, 481, 3, 2, 2, 2, 479, 477, 3, 2, 2, 2, 479,
	480, 3, 2, 2, 2, 480, 482, 3, 2, 2, 2, 481, 479, 3, 2, 2, 2, 482, 483,
	5, 88, 45, 2, 483, 484, 5, 90, 46, 2, 484, 85, 3, 2, 2, 2, 485, 486, 5,
	90, 46, 2, 486, 487, 7, 10, 2, 2, 487, 488, 5, 90, 46, 2, 488, 87, 3, 2,
	2, 2, 489, 490, 9, 4, 2, 2, 490, 89, 3, 2, 2, 2, 491, 492, 5, 92, 47, 2,
	492, 91, 3, 2, 2, 2, 493, 499, 5, 94, 48, 2, 494, 495, 7, 29, 2, 2, 495,
	496, 5, 90, 46, 2, 496, 497, 7, 7, 2, 2, 497, 498, 5, 90, 46, 2, 498, 500,
	3, 2, 2, 2, 499, 494, 3, 2, 2, 2, 499, 500, 3, 2, 2, 2, 500, 93, 3, 2,
	2, 2, 501, 502, 8, 48, 1, 2, 502, 503, 5, 96, 49, 2, 503, 509, 3, 2, 2,
	2, 504, 505, 12, 3, 2, 2, 505, 506, 7, 12, 2, 2, 506, 508, 5, 96, 49, 2,
	507, 504, 3, 2, 2, 2, 508, 511, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 509,
	510, 3, 2, 2, 2, 510, 95, 3, 2, 2, 2, 511, 509, 3, 2, 2, 2, 512, 513, 8,
	49, 1, 2, 513, 514, 5, 98, 50, 2, 514, 520, 3, 2, 2, 2, 515, 516, 12, 3,
	2, 2, 516, 517, 7, 13, 2, 2, 517, 519, 5, 98, 50, 2, 518, 515, 3, 2, 2,
	2, 519, 522, 3, 2, 2, 2, 520, 518, 3, 2, 2, 2, 520, 521, 3, 2, 2, 2, 521,
	97, 3, 2, 2, 2, 522, 520, 3, 2, 2, 2, 523, 524, 8, 50, 1, 2, 524, 525,
	5, 100, 51, 2, 525, 532, 3, 2, 2, 2, 526, 527, 12, 3, 2, 2, 527, 528, 5,
	120, 61, 2, 528, 529, 5, 100, 51, 2, 529, 531, 3, 2, 2, 2, 530, 526, 3,
	2, 2, 2, 531, 534, 3, 2, 2, 2, 532, 530, 3, 2, 2, 2, 532, 533, 3, 2, 2,
	2, 533, 99, 3, 2, 2, 2, 534, 532, 3, 2, 2, 2, 535, 536, 8, 51, 1, 2, 536,
	537, 5, 102, 52, 2, 537, 544, 3, 2, 2, 2, 538, 539, 12, 3, 2, 2, 539, 540,
	5, 122, 62, 2, 540, 541, 5, 102, 52, 2, 541, 543, 3, 2, 2, 2, 542, 538,
	3, 2, 2, 2, 543, 546, 3, 2, 2, 2, 544, 542, 3, 2, 2, 2, 544, 545, 3, 2,
	2, 2, 545, 101, 3, 2, 2, 2, 546, 544, 3, 2, 2, 2, 547, 550, 5, 104, 53,
	2, 548, 549, 7, 30, 2, 2, 549, 551, 5, 102, 52, 2, 550, 548, 3, 2, 2, 2,
	550, 551, 3, 2, 2, 2, 551, 103, 3, 2, 2, 2, 552, 553, 8, 53, 1, 2, 553,
	554, 5, 106, 54, 2, 554, 560, 3, 2, 2, 2, 555, 556, 12, 3, 2, 2, 556, 557,
	7, 32, 2, 2, 557, 559, 5, 38, 20, 2, 558, 555, 3, 2, 2, 2, 559, 562, 3,
	2, 2, 2, 560, 558, 3, 2, 2, 2, 560, 561, 3, 2, 2, 2, 561, 105, 3, 2, 2,
	2, 562, 560, 3, 2, 2, 2, 563, 564, 8, 54, 1, 2, 564, 565, 5, 108, 55, 2,
	565, 571, 3, 2, 2, 2, 566, 567, 12, 3, 2, 2, 567, 568, 7, 26, 2, 2, 568,
	570, 5, 108, 55, 2, 569, 566, 3, 2, 2, 2, 570, 573, 3, 2, 2, 2, 571, 569,
	3, 2, 2, 2, 571, 572, 3, 2, 2, 2, 572, 107, 3, 2, 2, 2, 573, 571, 3, 2,
	2, 2, 574, 575, 8, 55, 1, 2, 575, 576, 5, 110, 56, 2, 576, 583, 3, 2, 2,
	2, 577, 578, 12, 3, 2, 2, 578, 579, 5, 124, 63, 2, 579, 580, 5, 110, 56,
	2, 580, 582, 3, 2, 2, 2, 581, 577, 3, 2, 2, 2, 582, 585, 3, 2, 2, 2, 583,
	581, 3, 2, 2, 2, 583, 584, 3, 2, 2, 2, 584, 109, 3, 2, 2, 2, 585, 583,
	3, 2, 2, 2, 586, 587, 8, 56, 1, 2, 587, 588, 5, 112, 57, 2, 588, 595, 3,
	2, 2, 2, 589, 590, 12, 3, 2, 2, 590, 591, 5, 126, 64, 2, 591, 592, 5, 112,
	57, 2, 592, 594, 3, 2, 2, 2, 593, 589, 3, 2, 2, 2, 594, 597, 3, 2, 2, 2,
	595, 593, 3, 2, 2, 2, 595, 596, 3, 2, 2, 2, 596, 111, 3, 2, 2, 2, 597,
	595, 3, 2, 2, 2, 598, 607, 5, 114, 58, 2, 599, 601, 5, 128, 65, 2, 600,
	599, 3, 2, 2, 2, 601, 602, 3, 2, 2, 2, 602, 600, 3, 2, 2, 2, 602, 603,
	3, 2, 2, 2, 603, 604, 3, 2, 2, 2, 604, 605, 5, 112, 57, 2, 605, 607, 3,
	2, 2, 2, 606, 598, 3, 2, 2, 2, 606, 600, 3, 2, 2, 2, 607, 113, 3, 2, 2,
	2, 608, 613, 5, 132, 67, 2, 609, 613, 5, 134, 68, 2, 610, 613, 5, 136,
	69, 2, 611, 613, 5, 116, 59, 2, 612, 608, 3, 2, 2, 2, 612, 609, 3, 2, 2,
	2, 612, 610, 3, 2, 2, 2, 612, 611, 3, 2, 2, 2, 613, 115, 3, 2, 2, 2, 614,
	618, 5, 130, 66, 2, 615, 617, 5, 118, 60, 2, 616, 615, 3, 2, 2, 2, 617,
	620, 3, 2, 2, 2, 618, 616, 3, 2, 2, 2, 618, 619, 3, 2, 2, 2, 619, 117,
	3, 2, 2, 2, 620, 618, 3, 2, 2, 2, 621, 624, 5, 146, 74, 2, 622, 624, 5,
	152, 77, 2, 623, 621, 3, 2, 2, 2, 623, 622, 3, 2, 2, 2, 624, 119, 3, 2,
	2, 2, 625, 626, 9, 5, 2, 2, 626, 121, 3, 2, 2, 2, 627, 628, 9, 6, 2, 2,
	628, 123, 3, 2, 2, 2, 629, 630, 9, 7, 2, 2, 630, 125, 3, 2, 2, 2, 631,
	632, 9, 8, 2, 2, 632, 127, 3, 2, 2, 2, 633, 634, 9, 9, 2, 2, 634, 129,
	3, 2, 2, 2, 635, 640, 5, 138, 70, 2, 636, 640, 5, 140, 71, 2, 637, 640,
	5, 142, 72, 2, 638, 640, 5, 144, 73, 2, 639, 635, 3, 2, 2, 2, 639, 636,
	3, 2, 2, 2, 639, 637, 3, 2, 2, 2, 639, 638, 3, 2, 2, 2, 640, 131, 3, 2,
	2, 2, 641, 642, 7, 60, 2, 2, 642, 643, 5, 174, 88, 2, 643, 644, 5, 152,
	77, 2, 644, 133, 3, 2, 2, 2, 645, 646, 7, 61, 2, 2, 646, 647, 5, 90, 46,
	2, 647, 135, 3, 2, 2, 2, 648, 649, 7, 26, 2, 2, 649, 650, 5, 90, 46, 2,
	650, 651, 7, 31, 2, 2, 651, 652, 5, 40, 21, 2, 652, 137, 3, 2, 2, 2, 653,
	654, 5, 174, 88, 2, 654, 139, 3, 2, 2, 2, 655, 656, 5, 156, 79, 2, 656,
	141, 3, 2, 2, 2, 657, 658, 7, 40, 2, 2, 658, 661, 5, 34, 18, 2, 659, 660,
	7, 7, 2, 2, 660, 662, 5, 38, 20, 2, 661, 659, 3, 2, 2, 2, 661, 662, 3,
	2, 2, 2, 662, 663, 3, 2, 2, 2, 663, 664, 5, 56, 29, 2, 664, 143, 3, 2,
	2, 2, 665, 666, 7, 33, 2, 2, 666, 667, 5, 90, 46, 2, 667, 668, 7, 34, 2,
	2, 668, 145, 3, 2, 2, 2, 669, 672, 5, 148, 75, 2, 670, 672, 5, 150, 76,
	2, 671, 669, 3, 2, 2, 2, 671, 670, 3, 2, 2, 2, 672, 147, 3, 2, 2, 2, 673,
	674, 7, 14, 2, 2, 674, 675, 5, 174, 88, 2, 675, 149, 3, 2, 2, 2, 676, 679,
	7, 8, 2, 2, 677, 680, 5, 90, 46, 2, 678, 680, 5, 40, 21, 2, 679, 677, 3,
	2, 2, 2, 679, 678, 3, 2, 2, 2, 680, 681, 3, 2, 2, 2, 681, 682, 7, 9, 2,
	2, 682, 151, 3, 2, 2, 2, 683, 692, 7, 33, 2, 2, 684, 689, 5, 154, 78, 2,
	685, 686, 7, 4, 2, 2, 686, 688, 5, 154, 78, 2, 687, 685, 3, 2, 2, 2, 688,
	691, 3, 2, 2, 2, 689, 687, 3, 2, 2, 2, 689, 690, 3, 2, 2, 2, 690, 693,
	3, 2, 2, 2, 691, 689, 3, 2, 2, 2, 692, 684, 3, 2, 2, 2, 692, 693, 3, 2,
	2, 2, 693, 694, 3, 2, 2, 2, 694, 695, 7, 34, 2, 2, 695, 153, 3, 2, 2, 2,
	696, 697, 5, 174, 88, 2, 697, 698, 7, 7, 2, 2, 698, 700, 3, 2, 2, 2, 699,
	696, 3, 2, 2, 2, 699, 700, 3, 2, 2, 2, 700, 701, 3, 2, 2, 2, 701, 702,
	5, 90, 46, 2, 702, 155, 3, 2, 2, 2, 703, 710, 5, 164, 83, 2, 704, 710,
	5, 158, 80, 2, 705, 710, 5, 168, 85, 2, 706, 710, 5, 170, 86, 2, 707, 710,
	5, 162, 82, 2, 708, 710, 5, 160, 81, 2, 709, 703, 3, 2, 2, 2, 709, 704,
	3, 2, 2, 2, 709, 705, 3, 2, 2, 2, 709, 706, 3, 2, 2, 2, 709, 707, 3, 2,
	2, 2, 709, 708, 3, 2, 2, 2, 710, 157, 3, 2, 2, 2, 711, 712, 9, 10, 2, 2,
	712, 159, 3, 2, 2, 2, 713, 714, 7, 57, 2, 2, 714, 161, 3, 2, 2, 2, 715,
	716, 7, 68, 2, 2, 716, 163, 3, 2, 2, 2, 717, 719, 7, 22, 2, 2, 718, 717,
	3, 2, 2, 2, 718, 719, 3, 2, 2, 2, 719, 720, 3, 2, 2, 2, 720, 721, 5, 166,
	84, 2, 721, 165, 3, 2, 2, 2, 722, 728, 7, 63, 2, 2, 723, 728, 7, 64, 2,
	2, 724, 728, 7, 65, 2, 2, 725, 728, 7, 66, 2, 2, 726, 728, 7, 67, 2, 2,
	727, 722, 3, 2, 2, 2, 727, 723, 3, 2, 2, 2, 727, 724, 3, 2, 2, 2, 727,
	725, 3, 2, 2, 2, 727, 726, 3, 2, 2, 2, 728, 167, 3, 2, 2, 2, 729, 738,
	7, 8, 2, 2, 730, 735, 5, 90, 46, 2, 731, 732, 7, 4, 2, 2, 732, 734, 5,
	90, 46, 2, 733, 731, 3, 2, 2, 2, 734, 737, 3, 2, 2, 2, 735, 733, 3, 2,
	2, 2, 735, 736, 3, 2, 2, 2, 736, 739, 3, 2, 2, 2, 737, 735, 3, 2, 2, 2,
	738, 730, 3, 2, 2, 2, 738, 739, 3, 2, 2, 2, 739, 740, 3, 2, 2, 2, 740,
	741, 7, 9, 2, 2, 741, 169, 3, 2, 2, 2, 742, 751, 7, 5, 2, 2, 743, 748,
	5, 172, 87, 2, 744, 745, 7, 4, 2, 2, 745, 747, 5, 172, 87, 2, 746, 744,
	3, 2, 2, 2, 747, 750, 3, 2, 2, 2, 748, 746, 3, 2, 2, 2, 748, 749, 3, 2,
	2, 2, 749, 752, 3, 2, 2, 2, 750, 748, 3, 2, 2, 2, 751, 743, 3, 2, 2, 2,
	751, 752, 3, 2, 2, 2, 752, 753, 3, 2, 2, 2, 753, 754, 7, 6, 2, 2, 754,
	171, 3, 2, 2, 2, 755, 756, 5, 90, 46, 2, 756, 757, 7, 7, 2, 2, 757, 758,
	5, 90, 46, 2, 758, 173, 3, 2, 2, 2, 759, 760, 9, 11, 2, 2, 760, 175, 3,
	2, 2, 2, 761, 766, 7, 3, 2, 2, 762, 766, 7, 2, 2, 3, 763, 766, 6, 89, 15,
	2, 764, 766, 6, 89, 16, 2, 765, 761, 3, 2, 2, 2, 765, 762, 3, 2, 2, 2,
	765, 763, 3, 2, 2, 2, 765, 764, 3, 2, 2, 2, 766, 177, 3, 2, 2, 2, 69, 180,
	184, 191, 199, 207, 212, 216, 221, 237, 240, 246, 263, 271, 278, 288, 291,
	305, 308, 313, 320, 332, 335, 342, 353, 356, 385, 388, 408, 414, 421, 434,
	439, 448, 454, 456, 470, 479, 499, 509, 520, 532, 544, 550, 560, 571, 583,
	595, 602, 606, 612, 618, 623, 639, 661, 671, 679, 689, 692, 699, 709, 718,
	727, 735, 738, 748, 751, 765,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "','", "'{'", "'}'", "':'", "'['", "']'", "'<->'", "'='", "'||'",
	"'&&'", "'.'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'+'", "'-'",
	"'*'", "'/'", "'%'", "'&'", "'!'", "'<-'", "'?'", "", "'as'", "'as?'",
	"'('", "')'", "'transaction'", "'struct'", "'resource'", "'contract'",
	"'interface'", "'fun'", "'event'", "'emit'", "'pre'", "'post'", "'pub'",
	"'pub(set)'", "'return'", "'break'", "'continue'", "'let'", "'var'", "'if'",
	"'else'", "'while'", "'true'", "'false'", "'nil'", "'import'", "'from'",
	"'create'", "'destroy'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "Equal", "Unequal",
	"Less", "Greater", "LessEqual", "GreaterEqual", "Plus", "Minus", "Mul",
	"Div", "Mod", "Ampersand", "Negate", "Move", "Optional", "NilCoalescing",
	"Downcasting", "FailableDowncasting", "OpenParen", "CloseParen", "Transaction",
	"Struct", "Resource", "Contract", "Interface", "Fun", "Event", "Emit",
	"Pre", "Post", "Pub", "PubSet", "Return", "Break", "Continue", "Let", "Var",
	"If", "Else", "While", "True", "False", "Nil", "Import", "From", "Create",
	"Destroy", "Identifier", "DecimalLiteral", "BinaryLiteral", "OctalLiteral",
	"HexadecimalLiteral", "InvalidNumberLiteral", "StringLiteral", "WS", "Terminator",
	"BlockComment", "LineComment",
}

var ruleNames = []string{
	"program", "replInput", "declaration", "importDeclaration", "access", "compositeDeclaration",
	"conformances", "variableKind", "field", "interfaceDeclaration", "members",
	"member", "compositeKind", "specialFunctionDeclaration", "functionDeclaration",
	"eventDeclaration", "parameterList", "parameter", "typeAnnotation", "fullType",
	"baseType", "nominalType", "functionType", "variableSizedType", "constantSizedType",
	"dictionaryType", "block", "functionBlock", "preConditions", "postConditions",
	"conditions", "condition", "statements", "statement", "returnStatement",
	"breakStatement", "continueStatement", "ifStatement", "whileStatement",
	"emitStatement", "variableDeclaration", "assignment", "swap", "transfer",
	"expression", "conditionalExpression", "orExpression", "andExpression",
	"equalityExpression", "relationalExpression", "nilCoalescingExpression",
	"failableDowncastingExpression", "concatenatingExpression", "additiveExpression",
	"multiplicativeExpression", "unaryExpression", "primaryExpression", "composedExpression",
	"primaryExpressionSuffix", "equalityOp", "relationalOp", "additiveOp",
	"multiplicativeOp", "unaryOp", "primaryExpressionStart", "createExpression",
	"destroyExpression", "referenceExpression", "identifierExpression", "literalExpression",
	"functionExpression", "nestedExpression", "expressionAccess", "memberAccess",
	"bracketExpression", "invocation", "argument", "literal", "booleanLiteral",
	"nilLiteral", "stringLiteral", "integerLiteral", "positiveIntegerLiteral",
	"arrayLiteral", "dictionaryLiteral", "dictionaryEntry", "identifier", "eos",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type StrictusParser struct {
	*antlr.BaseParser
}

func NewStrictusParser(input antlr.TokenStream) *StrictusParser {
	this := new(StrictusParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Strictus.g4"

	return this
}

// Returns true if on the current index of the parser's
// token stream a token exists on the Hidden channel which
// either is a line terminator, or is a multi line comment that
// contains a line terminator.
func (p *StrictusParser) lineTerminatorAhead() bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	if ahead.GetChannel() != antlr.LexerHidden {
		// We're only interested in tokens on the HIDDEN channel.
		return true
	}

	if ahead.GetTokenType() == StrictusParserTerminator {
		// There is definitely a line terminator ahead.
		return true
	}

	if ahead.GetTokenType() == StrictusParserWS {
		// Get the token ahead of the current whitespaces.
		possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - 2
		ahead = p.GetTokenStream().Get(possibleIndexEosToken)
	}

	// Get the token's text and type.
	text := ahead.GetText()
	_type := ahead.GetTokenType()

	// Check if the token is, or contains a line terminator.
	return (_type == StrictusParserBlockComment && (strings.Contains(text, "\r") || strings.Contains(text, "\n"))) ||
		(_type == StrictusParserTerminator)
}

func (p *StrictusParser) noWhitespace() bool {
	index := p.GetCurrentToken().GetTokenIndex()
	return p.GetTokenStream().Get(index-1).GetTokenType() != StrictusParserWS
}

// StrictusParser tokens.
const (
	StrictusParserEOF                  = antlr.TokenEOF
	StrictusParserT__0                 = 1
	StrictusParserT__1                 = 2
	StrictusParserT__2                 = 3
	StrictusParserT__3                 = 4
	StrictusParserT__4                 = 5
	StrictusParserT__5                 = 6
	StrictusParserT__6                 = 7
	StrictusParserT__7                 = 8
	StrictusParserT__8                 = 9
	StrictusParserT__9                 = 10
	StrictusParserT__10                = 11
	StrictusParserT__11                = 12
	StrictusParserEqual                = 13
	StrictusParserUnequal              = 14
	StrictusParserLess                 = 15
	StrictusParserGreater              = 16
	StrictusParserLessEqual            = 17
	StrictusParserGreaterEqual         = 18
	StrictusParserPlus                 = 19
	StrictusParserMinus                = 20
	StrictusParserMul                  = 21
	StrictusParserDiv                  = 22
	StrictusParserMod                  = 23
	StrictusParserAmpersand            = 24
	StrictusParserNegate               = 25
	StrictusParserMove                 = 26
	StrictusParserOptional             = 27
	StrictusParserNilCoalescing        = 28
	StrictusParserDowncasting          = 29
	StrictusParserFailableDowncasting  = 30
	StrictusParserOpenParen            = 31
	StrictusParserCloseParen           = 32
	StrictusParserTransaction          = 33
	StrictusParserStruct               = 34
	StrictusParserResource             = 35
	StrictusParserContract             = 36
	StrictusParserInterface            = 37
	StrictusParserFun                  = 38
	StrictusParserEvent                = 39
	StrictusParserEmit                 = 40
	StrictusParserPre                  = 41
	StrictusParserPost                 = 42
	StrictusParserPub                  = 43
	StrictusParserPubSet               = 44
	StrictusParserReturn               = 45
	StrictusParserBreak                = 46
	StrictusParserContinue             = 47
	StrictusParserLet                  = 48
	StrictusParserVar                  = 49
	StrictusParserIf                   = 50
	StrictusParserElse                 = 51
	StrictusParserWhile                = 52
	StrictusParserTrue                 = 53
	StrictusParserFalse                = 54
	StrictusParserNil                  = 55
	StrictusParserImport               = 56
	StrictusParserFrom                 = 57
	StrictusParserCreate               = 58
	StrictusParserDestroy              = 59
	StrictusParserIdentifier           = 60
	StrictusParserDecimalLiteral       = 61
	StrictusParserBinaryLiteral        = 62
	StrictusParserOctalLiteral         = 63
	StrictusParserHexadecimalLiteral   = 64
	StrictusParserInvalidNumberLiteral = 65
	StrictusParserStringLiteral        = 66
	StrictusParserWS                   = 67
	StrictusParserTerminator           = 68
	StrictusParserBlockComment         = 69
	StrictusParserLineComment          = 70
)

// StrictusParser rules.
const (
	StrictusParserRULE_program                       = 0
	StrictusParserRULE_replInput                     = 1
	StrictusParserRULE_declaration                   = 2
	StrictusParserRULE_importDeclaration             = 3
	StrictusParserRULE_access                        = 4
	StrictusParserRULE_compositeDeclaration          = 5
	StrictusParserRULE_conformances                  = 6
	StrictusParserRULE_variableKind                  = 7
	StrictusParserRULE_field                         = 8
	StrictusParserRULE_interfaceDeclaration          = 9
	StrictusParserRULE_members                       = 10
	StrictusParserRULE_member                        = 11
	StrictusParserRULE_compositeKind                 = 12
	StrictusParserRULE_specialFunctionDeclaration    = 13
	StrictusParserRULE_functionDeclaration           = 14
	StrictusParserRULE_eventDeclaration              = 15
	StrictusParserRULE_parameterList                 = 16
	StrictusParserRULE_parameter                     = 17
	StrictusParserRULE_typeAnnotation                = 18
	StrictusParserRULE_fullType                      = 19
	StrictusParserRULE_baseType                      = 20
	StrictusParserRULE_nominalType                   = 21
	StrictusParserRULE_functionType                  = 22
	StrictusParserRULE_variableSizedType             = 23
	StrictusParserRULE_constantSizedType             = 24
	StrictusParserRULE_dictionaryType                = 25
	StrictusParserRULE_block                         = 26
	StrictusParserRULE_functionBlock                 = 27
	StrictusParserRULE_preConditions                 = 28
	StrictusParserRULE_postConditions                = 29
	StrictusParserRULE_conditions                    = 30
	StrictusParserRULE_condition                     = 31
	StrictusParserRULE_statements                    = 32
	StrictusParserRULE_statement                     = 33
	StrictusParserRULE_returnStatement               = 34
	StrictusParserRULE_breakStatement                = 35
	StrictusParserRULE_continueStatement             = 36
	StrictusParserRULE_ifStatement                   = 37
	StrictusParserRULE_whileStatement                = 38
	StrictusParserRULE_emitStatement                 = 39
	StrictusParserRULE_variableDeclaration           = 40
	StrictusParserRULE_assignment                    = 41
	StrictusParserRULE_swap                          = 42
	StrictusParserRULE_transfer                      = 43
	StrictusParserRULE_expression                    = 44
	StrictusParserRULE_conditionalExpression         = 45
	StrictusParserRULE_orExpression                  = 46
	StrictusParserRULE_andExpression                 = 47
	StrictusParserRULE_equalityExpression            = 48
	StrictusParserRULE_relationalExpression          = 49
	StrictusParserRULE_nilCoalescingExpression       = 50
	StrictusParserRULE_failableDowncastingExpression = 51
	StrictusParserRULE_concatenatingExpression       = 52
	StrictusParserRULE_additiveExpression            = 53
	StrictusParserRULE_multiplicativeExpression      = 54
	StrictusParserRULE_unaryExpression               = 55
	StrictusParserRULE_primaryExpression             = 56
	StrictusParserRULE_composedExpression            = 57
	StrictusParserRULE_primaryExpressionSuffix       = 58
	StrictusParserRULE_equalityOp                    = 59
	StrictusParserRULE_relationalOp                  = 60
	StrictusParserRULE_additiveOp                    = 61
	StrictusParserRULE_multiplicativeOp              = 62
	StrictusParserRULE_unaryOp                       = 63
	StrictusParserRULE_primaryExpressionStart        = 64
	StrictusParserRULE_createExpression              = 65
	StrictusParserRULE_destroyExpression             = 66
	StrictusParserRULE_referenceExpression           = 67
	StrictusParserRULE_identifierExpression          = 68
	StrictusParserRULE_literalExpression             = 69
	StrictusParserRULE_functionExpression            = 70
	StrictusParserRULE_nestedExpression              = 71
	StrictusParserRULE_expressionAccess              = 72
	StrictusParserRULE_memberAccess                  = 73
	StrictusParserRULE_bracketExpression             = 74
	StrictusParserRULE_invocation                    = 75
	StrictusParserRULE_argument                      = 76
	StrictusParserRULE_literal                       = 77
	StrictusParserRULE_booleanLiteral                = 78
	StrictusParserRULE_nilLiteral                    = 79
	StrictusParserRULE_stringLiteral                 = 80
	StrictusParserRULE_integerLiteral                = 81
	StrictusParserRULE_positiveIntegerLiteral        = 82
	StrictusParserRULE_arrayLiteral                  = 83
	StrictusParserRULE_dictionaryLiteral             = 84
	StrictusParserRULE_dictionaryEntry               = 85
	StrictusParserRULE_identifier                    = 86
	StrictusParserRULE_eos                           = 87
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(StrictusParserEOF, 0)
}

func (s *ProgramContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *ProgramContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StrictusParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(StrictusParserStruct-34))|(1<<(StrictusParserResource-34))|(1<<(StrictusParserContract-34))|(1<<(StrictusParserFun-34))|(1<<(StrictusParserEvent-34))|(1<<(StrictusParserPub-34))|(1<<(StrictusParserPubSet-34))|(1<<(StrictusParserLet-34))|(1<<(StrictusParserVar-34))|(1<<(StrictusParserImport-34)))) != 0 {
		{
			p.SetState(176)
			p.Declaration()
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StrictusParserT__0 {
			{
				p.SetState(177)
				p.Match(StrictusParserT__0)
			}

		}

		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(185)
		p.Match(StrictusParserEOF)
	}

	return localctx
}

// IReplInputContext is an interface to support dynamic dispatch.
type IReplInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReplInputContext differentiates from other interfaces.
	IsReplInputContext()
}

type ReplInputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplInputContext() *ReplInputContext {
	var p = new(ReplInputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_replInput
	return p
}

func (*ReplInputContext) IsReplInputContext() {}

func NewReplInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplInputContext {
	var p = new(ReplInputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_replInput

	return p
}

func (s *ReplInputContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplInputContext) Program() IProgramContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgramContext)
}

func (s *ReplInputContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ReplInputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplInputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplInputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterReplInput(s)
	}
}

func (s *ReplInputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitReplInput(s)
	}
}

func (s *ReplInputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitReplInput(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ReplInput() (localctx IReplInputContext) {
	localctx = NewReplInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StrictusParserRULE_replInput)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Program()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Statements()
		}

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) CompositeDeclaration() ICompositeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompositeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompositeDeclarationContext)
}

func (s *DeclarationContext) InterfaceDeclaration() IInterfaceDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceDeclarationContext)
}

func (s *DeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *DeclarationContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *DeclarationContext) ImportDeclaration() IImportDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportDeclarationContext)
}

func (s *DeclarationContext) EventDeclaration() IEventDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEventDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEventDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, StrictusParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.CompositeDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.InterfaceDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(193)
			p.FunctionDeclaration(true)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(194)
			p.VariableDeclaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(195)
			p.ImportDeclaration()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(196)
			p.EventDeclaration()
		}

	}

	return localctx
}

// IImportDeclarationContext is an interface to support dynamic dispatch.
type IImportDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclarationContext differentiates from other interfaces.
	IsImportDeclarationContext()
}

type ImportDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclarationContext() *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_importDeclaration
	return p
}

func (*ImportDeclarationContext) IsImportDeclarationContext() {}

func NewImportDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_importDeclaration

	return p
}

func (s *ImportDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclarationContext) Import() antlr.TerminalNode {
	return s.GetToken(StrictusParserImport, 0)
}

func (s *ImportDeclarationContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ImportDeclarationContext) HexadecimalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserHexadecimalLiteral, 0)
}

func (s *ImportDeclarationContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *ImportDeclarationContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ImportDeclarationContext) From() antlr.TerminalNode {
	return s.GetToken(StrictusParserFrom, 0)
}

func (s *ImportDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitImportDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ImportDeclaration() (localctx IImportDeclarationContext) {
	localctx = NewImportDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, StrictusParserRULE_importDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(StrictusParserImport)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(StrictusParserFrom-57))|(1<<(StrictusParserCreate-57))|(1<<(StrictusParserDestroy-57))|(1<<(StrictusParserIdentifier-57)))) != 0 {
		{
			p.SetState(200)
			p.Identifier()
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(201)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(202)
				p.Identifier()
			}

			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(208)
			p.Match(StrictusParserFrom)
		}

	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserStringLiteral:
		{
			p.SetState(212)
			p.StringLiteral()
		}

	case StrictusParserHexadecimalLiteral:
		{
			p.SetState(213)
			p.Match(StrictusParserHexadecimalLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAccessContext is an interface to support dynamic dispatch.
type IAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessContext differentiates from other interfaces.
	IsAccessContext()
}

type AccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessContext() *AccessContext {
	var p = new(AccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_access
	return p
}

func (*AccessContext) IsAccessContext() {}

func NewAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessContext {
	var p = new(AccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_access

	return p
}

func (s *AccessContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessContext) Pub() antlr.TerminalNode {
	return s.GetToken(StrictusParserPub, 0)
}

func (s *AccessContext) PubSet() antlr.TerminalNode {
	return s.GetToken(StrictusParserPubSet, 0)
}

func (s *AccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAccess(s)
	}
}

func (s *AccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAccess(s)
	}
}

func (s *AccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Access() (localctx IAccessContext) {
	localctx = NewAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, StrictusParserRULE_access)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserStruct, StrictusParserResource, StrictusParserContract, StrictusParserFun, StrictusParserLet, StrictusParserVar, StrictusParserFrom, StrictusParserCreate, StrictusParserDestroy, StrictusParserIdentifier:
		p.EnterOuterAlt(localctx, 1)

	case StrictusParserPub:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(StrictusParserPub)
		}

	case StrictusParserPubSet:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(218)
			p.Match(StrictusParserPubSet)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICompositeDeclarationContext is an interface to support dynamic dispatch.
type ICompositeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompositeDeclarationContext differentiates from other interfaces.
	IsCompositeDeclarationContext()
}

type CompositeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompositeDeclarationContext() *CompositeDeclarationContext {
	var p = new(CompositeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_compositeDeclaration
	return p
}

func (*CompositeDeclarationContext) IsCompositeDeclarationContext() {}

func NewCompositeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompositeDeclarationContext {
	var p = new(CompositeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_compositeDeclaration

	return p
}

func (s *CompositeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *CompositeDeclarationContext) Access() IAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *CompositeDeclarationContext) CompositeKind() ICompositeKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompositeKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompositeKindContext)
}

func (s *CompositeDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *CompositeDeclarationContext) Conformances() IConformancesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConformancesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConformancesContext)
}

func (s *CompositeDeclarationContext) Members() IMembersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMembersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMembersContext)
}

func (s *CompositeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompositeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompositeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterCompositeDeclaration(s)
	}
}

func (s *CompositeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitCompositeDeclaration(s)
	}
}

func (s *CompositeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitCompositeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) CompositeDeclaration() (localctx ICompositeDeclarationContext) {
	localctx = NewCompositeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, StrictusParserRULE_compositeDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Access()
	}
	{
		p.SetState(222)
		p.CompositeKind()
	}
	{
		p.SetState(223)
		p.Identifier()
	}
	{
		p.SetState(224)
		p.Conformances()
	}
	{
		p.SetState(225)
		p.Match(StrictusParserT__2)
	}
	{
		p.SetState(226)
		p.Members(true)
	}
	{
		p.SetState(227)
		p.Match(StrictusParserT__3)
	}

	return localctx
}

// IConformancesContext is an interface to support dynamic dispatch.
type IConformancesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConformancesContext differentiates from other interfaces.
	IsConformancesContext()
}

type ConformancesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConformancesContext() *ConformancesContext {
	var p = new(ConformancesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_conformances
	return p
}

func (*ConformancesContext) IsConformancesContext() {}

func NewConformancesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConformancesContext {
	var p = new(ConformancesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_conformances

	return p
}

func (s *ConformancesContext) GetParser() antlr.Parser { return s.parser }

func (s *ConformancesContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *ConformancesContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ConformancesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConformancesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConformancesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterConformances(s)
	}
}

func (s *ConformancesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitConformances(s)
	}
}

func (s *ConformancesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitConformances(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Conformances() (localctx IConformancesContext) {
	localctx = NewConformancesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, StrictusParserRULE_conformances)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserT__4 {
		{
			p.SetState(229)
			p.Match(StrictusParserT__4)
		}
		{
			p.SetState(230)
			p.Identifier()
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(231)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(232)
				p.Identifier()
			}

			p.SetState(237)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IVariableKindContext is an interface to support dynamic dispatch.
type IVariableKindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableKindContext differentiates from other interfaces.
	IsVariableKindContext()
}

type VariableKindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableKindContext() *VariableKindContext {
	var p = new(VariableKindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_variableKind
	return p
}

func (*VariableKindContext) IsVariableKindContext() {}

func NewVariableKindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableKindContext {
	var p = new(VariableKindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_variableKind

	return p
}

func (s *VariableKindContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableKindContext) Let() antlr.TerminalNode {
	return s.GetToken(StrictusParserLet, 0)
}

func (s *VariableKindContext) Var() antlr.TerminalNode {
	return s.GetToken(StrictusParserVar, 0)
}

func (s *VariableKindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableKindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableKindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterVariableKind(s)
	}
}

func (s *VariableKindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitVariableKind(s)
	}
}

func (s *VariableKindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitVariableKind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) VariableKind() (localctx IVariableKindContext) {
	localctx = NewVariableKindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, StrictusParserRULE_variableKind)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserLet || _la == StrictusParserVar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Access() IAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *FieldContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FieldContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *FieldContext) VariableKind() IVariableKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableKindContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, StrictusParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Access()
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserLet || _la == StrictusParserVar {
		{
			p.SetState(243)
			p.VariableKind()
		}

	}
	{
		p.SetState(246)
		p.Identifier()
	}
	{
		p.SetState(247)
		p.Match(StrictusParserT__4)
	}
	{
		p.SetState(248)
		p.TypeAnnotation()
	}

	return localctx
}

// IInterfaceDeclarationContext is an interface to support dynamic dispatch.
type IInterfaceDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceDeclarationContext differentiates from other interfaces.
	IsInterfaceDeclarationContext()
}

type InterfaceDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceDeclarationContext() *InterfaceDeclarationContext {
	var p = new(InterfaceDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_interfaceDeclaration
	return p
}

func (*InterfaceDeclarationContext) IsInterfaceDeclarationContext() {}

func NewInterfaceDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceDeclarationContext {
	var p = new(InterfaceDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_interfaceDeclaration

	return p
}

func (s *InterfaceDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceDeclarationContext) Access() IAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *InterfaceDeclarationContext) CompositeKind() ICompositeKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompositeKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompositeKindContext)
}

func (s *InterfaceDeclarationContext) Interface() antlr.TerminalNode {
	return s.GetToken(StrictusParserInterface, 0)
}

func (s *InterfaceDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *InterfaceDeclarationContext) Members() IMembersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMembersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMembersContext)
}

func (s *InterfaceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterInterfaceDeclaration(s)
	}
}

func (s *InterfaceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitInterfaceDeclaration(s)
	}
}

func (s *InterfaceDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitInterfaceDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) InterfaceDeclaration() (localctx IInterfaceDeclarationContext) {
	localctx = NewInterfaceDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, StrictusParserRULE_interfaceDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Access()
	}
	{
		p.SetState(251)
		p.CompositeKind()
	}
	{
		p.SetState(252)
		p.Match(StrictusParserInterface)
	}
	{
		p.SetState(253)
		p.Identifier()
	}
	{
		p.SetState(254)
		p.Match(StrictusParserT__2)
	}
	{
		p.SetState(255)
		p.Members(false)
	}
	{
		p.SetState(256)
		p.Match(StrictusParserT__3)
	}

	return localctx
}

// IMembersContext is an interface to support dynamic dispatch.
type IMembersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFunctionBlockRequired returns the functionBlockRequired attribute.
	GetFunctionBlockRequired() bool

	// SetFunctionBlockRequired sets the functionBlockRequired attribute.
	SetFunctionBlockRequired(bool)

	// IsMembersContext differentiates from other interfaces.
	IsMembersContext()
}

type MembersContext struct {
	*antlr.BaseParserRuleContext
	parser                antlr.Parser
	functionBlockRequired bool
}

func NewEmptyMembersContext() *MembersContext {
	var p = new(MembersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_members
	return p
}

func (*MembersContext) IsMembersContext() {}

func NewMembersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, functionBlockRequired bool) *MembersContext {
	var p = new(MembersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_members

	p.functionBlockRequired = functionBlockRequired

	return p
}

func (s *MembersContext) GetParser() antlr.Parser { return s.parser }

func (s *MembersContext) GetFunctionBlockRequired() bool { return s.functionBlockRequired }

func (s *MembersContext) SetFunctionBlockRequired(v bool) { s.functionBlockRequired = v }

func (s *MembersContext) AllMember() []IMemberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemberContext)(nil)).Elem())
	var tst = make([]IMemberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemberContext)
		}
	}

	return tst
}

func (s *MembersContext) Member(i int) IMemberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemberContext)
}

func (s *MembersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MembersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterMembers(s)
	}
}

func (s *MembersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitMembers(s)
	}
}

func (s *MembersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitMembers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Members(functionBlockRequired bool) (localctx IMembersContext) {
	localctx = NewMembersContext(p, p.GetParserRuleContext(), p.GetState(), functionBlockRequired)
	p.EnterRule(localctx, 20, StrictusParserRULE_members)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(StrictusParserStruct-34))|(1<<(StrictusParserResource-34))|(1<<(StrictusParserContract-34))|(1<<(StrictusParserFun-34))|(1<<(StrictusParserPub-34))|(1<<(StrictusParserPubSet-34))|(1<<(StrictusParserLet-34))|(1<<(StrictusParserVar-34))|(1<<(StrictusParserFrom-34))|(1<<(StrictusParserCreate-34))|(1<<(StrictusParserDestroy-34))|(1<<(StrictusParserIdentifier-34)))) != 0 {
		{
			p.SetState(258)
			p.Member(functionBlockRequired)
		}

		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMemberContext is an interface to support dynamic dispatch.
type IMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFunctionBlockRequired returns the functionBlockRequired attribute.
	GetFunctionBlockRequired() bool

	// SetFunctionBlockRequired sets the functionBlockRequired attribute.
	SetFunctionBlockRequired(bool)

	// IsMemberContext differentiates from other interfaces.
	IsMemberContext()
}

type MemberContext struct {
	*antlr.BaseParserRuleContext
	parser                antlr.Parser
	functionBlockRequired bool
}

func NewEmptyMemberContext() *MemberContext {
	var p = new(MemberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_member
	return p
}

func (*MemberContext) IsMemberContext() {}

func NewMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, functionBlockRequired bool) *MemberContext {
	var p = new(MemberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_member

	p.functionBlockRequired = functionBlockRequired

	return p
}

func (s *MemberContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberContext) GetFunctionBlockRequired() bool { return s.functionBlockRequired }

func (s *MemberContext) SetFunctionBlockRequired(v bool) { s.functionBlockRequired = v }

func (s *MemberContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MemberContext) SpecialFunctionDeclaration() ISpecialFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecialFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecialFunctionDeclarationContext)
}

func (s *MemberContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *MemberContext) InterfaceDeclaration() IInterfaceDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceDeclarationContext)
}

func (s *MemberContext) CompositeDeclaration() ICompositeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompositeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompositeDeclarationContext)
}

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterMember(s)
	}
}

func (s *MemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitMember(s)
	}
}

func (s *MemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitMember(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Member(functionBlockRequired bool) (localctx IMemberContext) {
	localctx = NewMemberContext(p, p.GetParserRuleContext(), p.GetState(), functionBlockRequired)
	p.EnterRule(localctx, 22, StrictusParserRULE_member)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)
			p.SpecialFunctionDeclaration(functionBlockRequired)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(266)
			p.FunctionDeclaration(functionBlockRequired)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(267)
			p.InterfaceDeclaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(268)
			p.CompositeDeclaration()
		}

	}

	return localctx
}

// ICompositeKindContext is an interface to support dynamic dispatch.
type ICompositeKindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompositeKindContext differentiates from other interfaces.
	IsCompositeKindContext()
}

type CompositeKindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompositeKindContext() *CompositeKindContext {
	var p = new(CompositeKindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_compositeKind
	return p
}

func (*CompositeKindContext) IsCompositeKindContext() {}

func NewCompositeKindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompositeKindContext {
	var p = new(CompositeKindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_compositeKind

	return p
}

func (s *CompositeKindContext) GetParser() antlr.Parser { return s.parser }

func (s *CompositeKindContext) Struct() antlr.TerminalNode {
	return s.GetToken(StrictusParserStruct, 0)
}

func (s *CompositeKindContext) Resource() antlr.TerminalNode {
	return s.GetToken(StrictusParserResource, 0)
}

func (s *CompositeKindContext) Contract() antlr.TerminalNode {
	return s.GetToken(StrictusParserContract, 0)
}

func (s *CompositeKindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompositeKindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompositeKindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterCompositeKind(s)
	}
}

func (s *CompositeKindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitCompositeKind(s)
	}
}

func (s *CompositeKindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitCompositeKind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) CompositeKind() (localctx ICompositeKindContext) {
	localctx = NewCompositeKindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, StrictusParserRULE_compositeKind)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(StrictusParserStruct-34))|(1<<(StrictusParserResource-34))|(1<<(StrictusParserContract-34)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISpecialFunctionDeclarationContext is an interface to support dynamic dispatch.
type ISpecialFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetB returns the b rule contexts.
	GetB() IFunctionBlockContext

	// SetB sets the b rule contexts.
	SetB(IFunctionBlockContext)

	// GetFunctionBlockRequired returns the functionBlockRequired attribute.
	GetFunctionBlockRequired() bool

	// SetFunctionBlockRequired sets the functionBlockRequired attribute.
	SetFunctionBlockRequired(bool)

	// IsSpecialFunctionDeclarationContext differentiates from other interfaces.
	IsSpecialFunctionDeclarationContext()
}

type SpecialFunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser                antlr.Parser
	functionBlockRequired bool
	b                     IFunctionBlockContext
}

func NewEmptySpecialFunctionDeclarationContext() *SpecialFunctionDeclarationContext {
	var p = new(SpecialFunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_specialFunctionDeclaration
	return p
}

func (*SpecialFunctionDeclarationContext) IsSpecialFunctionDeclarationContext() {}

func NewSpecialFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, functionBlockRequired bool) *SpecialFunctionDeclarationContext {
	var p = new(SpecialFunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_specialFunctionDeclaration

	p.functionBlockRequired = functionBlockRequired

	return p
}

func (s *SpecialFunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecialFunctionDeclarationContext) GetB() IFunctionBlockContext { return s.b }

func (s *SpecialFunctionDeclarationContext) SetB(v IFunctionBlockContext) { s.b = v }

func (s *SpecialFunctionDeclarationContext) GetFunctionBlockRequired() bool {
	return s.functionBlockRequired
}

func (s *SpecialFunctionDeclarationContext) SetFunctionBlockRequired(v bool) {
	s.functionBlockRequired = v
}

func (s *SpecialFunctionDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SpecialFunctionDeclarationContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *SpecialFunctionDeclarationContext) FunctionBlock() IFunctionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBlockContext)
}

func (s *SpecialFunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecialFunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecialFunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterSpecialFunctionDeclaration(s)
	}
}

func (s *SpecialFunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitSpecialFunctionDeclaration(s)
	}
}

func (s *SpecialFunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitSpecialFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) SpecialFunctionDeclaration(functionBlockRequired bool) (localctx ISpecialFunctionDeclarationContext) {
	localctx = NewSpecialFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState(), functionBlockRequired)
	p.EnterRule(localctx, 26, StrictusParserRULE_specialFunctionDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Identifier()
	}
	{
		p.SetState(274)
		p.ParameterList()
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(275)

			var _x = p.FunctionBlock()

			localctx.(*SpecialFunctionDeclarationContext).b = _x
		}

	}
	p.SetState(278)

	if !(!localctx.(*SpecialFunctionDeclarationContext).functionBlockRequired || localctx.(*SpecialFunctionDeclarationContext).b != nil) {
		panic(antlr.NewFailedPredicateException(p, " !$functionBlockRequired || $ctx.b != nil ", ""))
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeAnnotationContext

	// GetB returns the b rule contexts.
	GetB() IFunctionBlockContext

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeAnnotationContext)

	// SetB sets the b rule contexts.
	SetB(IFunctionBlockContext)

	// GetFunctionBlockRequired returns the functionBlockRequired attribute.
	GetFunctionBlockRequired() bool

	// SetFunctionBlockRequired sets the functionBlockRequired attribute.
	SetFunctionBlockRequired(bool)

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser                antlr.Parser
	functionBlockRequired bool
	returnType            ITypeAnnotationContext
	b                     IFunctionBlockContext
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, functionBlockRequired bool) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_functionDeclaration

	p.functionBlockRequired = functionBlockRequired

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) GetReturnType() ITypeAnnotationContext { return s.returnType }

func (s *FunctionDeclarationContext) GetB() IFunctionBlockContext { return s.b }

func (s *FunctionDeclarationContext) SetReturnType(v ITypeAnnotationContext) { s.returnType = v }

func (s *FunctionDeclarationContext) SetB(v IFunctionBlockContext) { s.b = v }

func (s *FunctionDeclarationContext) GetFunctionBlockRequired() bool { return s.functionBlockRequired }

func (s *FunctionDeclarationContext) SetFunctionBlockRequired(v bool) { s.functionBlockRequired = v }

func (s *FunctionDeclarationContext) Access() IAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *FunctionDeclarationContext) Fun() antlr.TerminalNode {
	return s.GetToken(StrictusParserFun, 0)
}

func (s *FunctionDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FunctionDeclarationContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionDeclarationContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *FunctionDeclarationContext) FunctionBlock() IFunctionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBlockContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) FunctionDeclaration(functionBlockRequired bool) (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState(), functionBlockRequired)
	p.EnterRule(localctx, 28, StrictusParserRULE_functionDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Access()
	}
	{
		p.SetState(281)
		p.Match(StrictusParserFun)
	}
	{
		p.SetState(282)
		p.Identifier()
	}
	{
		p.SetState(283)
		p.ParameterList()
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(284)
			p.Match(StrictusParserT__4)
		}
		{
			p.SetState(285)

			var _x = p.TypeAnnotation()

			localctx.(*FunctionDeclarationContext).returnType = _x
		}

	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(288)

			var _x = p.FunctionBlock()

			localctx.(*FunctionDeclarationContext).b = _x
		}

	}
	p.SetState(291)

	if !(!localctx.(*FunctionDeclarationContext).functionBlockRequired || localctx.(*FunctionDeclarationContext).b != nil) {
		panic(antlr.NewFailedPredicateException(p, " !$functionBlockRequired || $ctx.b != nil ", ""))
	}

	return localctx
}

// IEventDeclarationContext is an interface to support dynamic dispatch.
type IEventDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEventDeclarationContext differentiates from other interfaces.
	IsEventDeclarationContext()
}

type EventDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEventDeclarationContext() *EventDeclarationContext {
	var p = new(EventDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_eventDeclaration
	return p
}

func (*EventDeclarationContext) IsEventDeclarationContext() {}

func NewEventDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventDeclarationContext {
	var p = new(EventDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_eventDeclaration

	return p
}

func (s *EventDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *EventDeclarationContext) Event() antlr.TerminalNode {
	return s.GetToken(StrictusParserEvent, 0)
}

func (s *EventDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *EventDeclarationContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *EventDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterEventDeclaration(s)
	}
}

func (s *EventDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitEventDeclaration(s)
	}
}

func (s *EventDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitEventDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) EventDeclaration() (localctx IEventDeclarationContext) {
	localctx = NewEventDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, StrictusParserRULE_eventDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(StrictusParserEvent)
	}
	{
		p.SetState(294)
		p.Identifier()
	}
	{
		p.SetState(295)
		p.ParameterList()
	}

	return localctx
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_parameterList
	return p
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, 0)
}

func (s *ParameterListContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, 0)
}

func (s *ParameterListContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, StrictusParserRULE_parameterList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(StrictusParserOpenParen)
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(StrictusParserFrom-57))|(1<<(StrictusParserCreate-57))|(1<<(StrictusParserDestroy-57))|(1<<(StrictusParserIdentifier-57)))) != 0 {
		{
			p.SetState(298)
			p.Parameter()
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(299)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(300)
				p.Parameter()
			}

			p.SetState(305)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(308)
		p.Match(StrictusParserCloseParen)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArgumentLabel returns the argumentLabel rule contexts.
	GetArgumentLabel() IIdentifierContext

	// GetParameterName returns the parameterName rule contexts.
	GetParameterName() IIdentifierContext

	// SetArgumentLabel sets the argumentLabel rule contexts.
	SetArgumentLabel(IIdentifierContext)

	// SetParameterName sets the parameterName rule contexts.
	SetParameterName(IIdentifierContext)

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	argumentLabel IIdentifierContext
	parameterName IIdentifierContext
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) GetArgumentLabel() IIdentifierContext { return s.argumentLabel }

func (s *ParameterContext) GetParameterName() IIdentifierContext { return s.parameterName }

func (s *ParameterContext) SetArgumentLabel(v IIdentifierContext) { s.argumentLabel = v }

func (s *ParameterContext) SetParameterName(v IIdentifierContext) { s.parameterName = v }

func (s *ParameterContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *ParameterContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *ParameterContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, StrictusParserRULE_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(311)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(310)

			var _x = p.Identifier()

			localctx.(*ParameterContext).argumentLabel = _x
		}

	}
	{
		p.SetState(313)

		var _x = p.Identifier()

		localctx.(*ParameterContext).parameterName = _x
	}
	{
		p.SetState(314)
		p.Match(StrictusParserT__4)
	}
	{
		p.SetState(315)
		p.TypeAnnotation()
	}

	return localctx
}

// ITypeAnnotationContext is an interface to support dynamic dispatch.
type ITypeAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeAnnotationContext differentiates from other interfaces.
	IsTypeAnnotationContext()
}

type TypeAnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAnnotationContext() *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_typeAnnotation
	return p
}

func (*TypeAnnotationContext) IsTypeAnnotationContext() {}

func NewTypeAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_typeAnnotation

	return p
}

func (s *TypeAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAnnotationContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *TypeAnnotationContext) Move() antlr.TerminalNode {
	return s.GetToken(StrictusParserMove, 0)
}

func (s *TypeAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitTypeAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) TypeAnnotation() (localctx ITypeAnnotationContext) {
	localctx = NewTypeAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, StrictusParserRULE_typeAnnotation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserMove {
		{
			p.SetState(317)
			p.Match(StrictusParserMove)
		}

	}
	{
		p.SetState(320)
		p.FullType()
	}

	return localctx
}

// IFullTypeContext is an interface to support dynamic dispatch.
type IFullTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReference returns the reference token.
	GetReference() antlr.Token

	// Get_Optional returns the _Optional token.
	Get_Optional() antlr.Token

	// SetReference sets the reference token.
	SetReference(antlr.Token)

	// Set_Optional sets the _Optional token.
	Set_Optional(antlr.Token)

	// GetOptionals returns the optionals token list.
	GetOptionals() []antlr.Token

	// SetOptionals sets the optionals token list.
	SetOptionals([]antlr.Token)

	// IsFullTypeContext differentiates from other interfaces.
	IsFullTypeContext()
}

type FullTypeContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	reference antlr.Token
	_Optional antlr.Token
	optionals []antlr.Token
}

func NewEmptyFullTypeContext() *FullTypeContext {
	var p = new(FullTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_fullType
	return p
}

func (*FullTypeContext) IsFullTypeContext() {}

func NewFullTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullTypeContext {
	var p = new(FullTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_fullType

	return p
}

func (s *FullTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FullTypeContext) GetReference() antlr.Token { return s.reference }

func (s *FullTypeContext) Get_Optional() antlr.Token { return s._Optional }

func (s *FullTypeContext) SetReference(v antlr.Token) { s.reference = v }

func (s *FullTypeContext) Set_Optional(v antlr.Token) { s._Optional = v }

func (s *FullTypeContext) GetOptionals() []antlr.Token { return s.optionals }

func (s *FullTypeContext) SetOptionals(v []antlr.Token) { s.optionals = v }

func (s *FullTypeContext) BaseType() IBaseTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *FullTypeContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(StrictusParserAmpersand, 0)
}

func (s *FullTypeContext) AllOptional() []antlr.TerminalNode {
	return s.GetTokens(StrictusParserOptional)
}

func (s *FullTypeContext) Optional(i int) antlr.TerminalNode {
	return s.GetToken(StrictusParserOptional, i)
}

func (s *FullTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFullType(s)
	}
}

func (s *FullTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFullType(s)
	}
}

func (s *FullTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFullType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) FullType() (localctx IFullTypeContext) {
	localctx = NewFullTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, StrictusParserRULE_fullType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(333)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserAmpersand:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)

			var _m = p.Match(StrictusParserAmpersand)

			localctx.(*FullTypeContext).reference = _m
		}
		p.SetState(323)

		if !(p.noWhitespace()) {
			panic(antlr.NewFailedPredicateException(p, "p.noWhitespace()", ""))
		}
		{
			p.SetState(324)
			p.BaseType()
		}

	case StrictusParserT__2, StrictusParserT__5, StrictusParserOpenParen, StrictusParserFrom, StrictusParserCreate, StrictusParserDestroy, StrictusParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(325)
			p.BaseType()
		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(326)

				if !(p.noWhitespace()) {
					panic(antlr.NewFailedPredicateException(p, "p.noWhitespace()", ""))
				}
				{
					p.SetState(327)

					var _m = p.Match(StrictusParserOptional)

					localctx.(*FullTypeContext)._Optional = _m
				}
				localctx.(*FullTypeContext).optionals = append(localctx.(*FullTypeContext).optionals, localctx.(*FullTypeContext)._Optional)

			}
			p.SetState(332)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBaseTypeContext is an interface to support dynamic dispatch.
type IBaseTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseTypeContext differentiates from other interfaces.
	IsBaseTypeContext()
}

type BaseTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseTypeContext() *BaseTypeContext {
	var p = new(BaseTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_baseType
	return p
}

func (*BaseTypeContext) IsBaseTypeContext() {}

func NewBaseTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseTypeContext {
	var p = new(BaseTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_baseType

	return p
}

func (s *BaseTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseTypeContext) NominalType() INominalTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INominalTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INominalTypeContext)
}

func (s *BaseTypeContext) FunctionType() IFunctionTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionTypeContext)
}

func (s *BaseTypeContext) VariableSizedType() IVariableSizedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableSizedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableSizedTypeContext)
}

func (s *BaseTypeContext) ConstantSizedType() IConstantSizedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantSizedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantSizedTypeContext)
}

func (s *BaseTypeContext) DictionaryType() IDictionaryTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictionaryTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDictionaryTypeContext)
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (s *BaseTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBaseType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) BaseType() (localctx IBaseTypeContext) {
	localctx = NewBaseTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, StrictusParserRULE_baseType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(335)
			p.NominalType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(336)
			p.FunctionType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(337)
			p.VariableSizedType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(338)
			p.ConstantSizedType()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(339)
			p.DictionaryType()
		}

	}

	return localctx
}

// INominalTypeContext is an interface to support dynamic dispatch.
type INominalTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNominalTypeContext differentiates from other interfaces.
	IsNominalTypeContext()
}

type NominalTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNominalTypeContext() *NominalTypeContext {
	var p = new(NominalTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_nominalType
	return p
}

func (*NominalTypeContext) IsNominalTypeContext() {}

func NewNominalTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NominalTypeContext {
	var p = new(NominalTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_nominalType

	return p
}

func (s *NominalTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NominalTypeContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *NominalTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NominalTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NominalTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterNominalType(s)
	}
}

func (s *NominalTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitNominalType(s)
	}
}

func (s *NominalTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitNominalType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) NominalType() (localctx INominalTypeContext) {
	localctx = NewNominalTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, StrictusParserRULE_nominalType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Identifier()
	}

	return localctx
}

// IFunctionTypeContext is an interface to support dynamic dispatch.
type IFunctionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_typeAnnotation returns the _typeAnnotation rule contexts.
	Get_typeAnnotation() ITypeAnnotationContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeAnnotationContext

	// Set_typeAnnotation sets the _typeAnnotation rule contexts.
	Set_typeAnnotation(ITypeAnnotationContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeAnnotationContext)

	// GetParameterTypes returns the parameterTypes rule context list.
	GetParameterTypes() []ITypeAnnotationContext

	// SetParameterTypes sets the parameterTypes rule context list.
	SetParameterTypes([]ITypeAnnotationContext)

	// IsFunctionTypeContext differentiates from other interfaces.
	IsFunctionTypeContext()
}

type FunctionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	_typeAnnotation ITypeAnnotationContext
	parameterTypes  []ITypeAnnotationContext
	returnType      ITypeAnnotationContext
}

func NewEmptyFunctionTypeContext() *FunctionTypeContext {
	var p = new(FunctionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_functionType
	return p
}

func (*FunctionTypeContext) IsFunctionTypeContext() {}

func NewFunctionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionTypeContext {
	var p = new(FunctionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_functionType

	return p
}

func (s *FunctionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionTypeContext) Get_typeAnnotation() ITypeAnnotationContext { return s._typeAnnotation }

func (s *FunctionTypeContext) GetReturnType() ITypeAnnotationContext { return s.returnType }

func (s *FunctionTypeContext) Set_typeAnnotation(v ITypeAnnotationContext) { s._typeAnnotation = v }

func (s *FunctionTypeContext) SetReturnType(v ITypeAnnotationContext) { s.returnType = v }

func (s *FunctionTypeContext) GetParameterTypes() []ITypeAnnotationContext { return s.parameterTypes }

func (s *FunctionTypeContext) SetParameterTypes(v []ITypeAnnotationContext) { s.parameterTypes = v }

func (s *FunctionTypeContext) AllOpenParen() []antlr.TerminalNode {
	return s.GetTokens(StrictusParserOpenParen)
}

func (s *FunctionTypeContext) OpenParen(i int) antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, i)
}

func (s *FunctionTypeContext) AllCloseParen() []antlr.TerminalNode {
	return s.GetTokens(StrictusParserCloseParen)
}

func (s *FunctionTypeContext) CloseParen(i int) antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, i)
}

func (s *FunctionTypeContext) AllTypeAnnotation() []ITypeAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem())
	var tst = make([]ITypeAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeAnnotationContext)
		}
	}

	return tst
}

func (s *FunctionTypeContext) TypeAnnotation(i int) ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *FunctionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFunctionType(s)
	}
}

func (s *FunctionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFunctionType(s)
	}
}

func (s *FunctionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFunctionType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) FunctionType() (localctx IFunctionTypeContext) {
	localctx = NewFunctionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, StrictusParserRULE_functionType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(StrictusParserOpenParen)
	}
	{
		p.SetState(345)
		p.Match(StrictusParserOpenParen)
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__2)|(1<<StrictusParserT__5)|(1<<StrictusParserAmpersand)|(1<<StrictusParserMove)|(1<<StrictusParserOpenParen))) != 0) || (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(StrictusParserFrom-57))|(1<<(StrictusParserCreate-57))|(1<<(StrictusParserDestroy-57))|(1<<(StrictusParserIdentifier-57)))) != 0) {
		{
			p.SetState(346)

			var _x = p.TypeAnnotation()

			localctx.(*FunctionTypeContext)._typeAnnotation = _x
		}
		localctx.(*FunctionTypeContext).parameterTypes = append(localctx.(*FunctionTypeContext).parameterTypes, localctx.(*FunctionTypeContext)._typeAnnotation)
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(347)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(348)

				var _x = p.TypeAnnotation()

				localctx.(*FunctionTypeContext)._typeAnnotation = _x
			}
			localctx.(*FunctionTypeContext).parameterTypes = append(localctx.(*FunctionTypeContext).parameterTypes, localctx.(*FunctionTypeContext)._typeAnnotation)

			p.SetState(353)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(356)
		p.Match(StrictusParserCloseParen)
	}
	{
		p.SetState(357)
		p.Match(StrictusParserT__4)
	}
	{
		p.SetState(358)

		var _x = p.TypeAnnotation()

		localctx.(*FunctionTypeContext).returnType = _x
	}
	{
		p.SetState(359)
		p.Match(StrictusParserCloseParen)
	}

	return localctx
}

// IVariableSizedTypeContext is an interface to support dynamic dispatch.
type IVariableSizedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableSizedTypeContext differentiates from other interfaces.
	IsVariableSizedTypeContext()
}

type VariableSizedTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableSizedTypeContext() *VariableSizedTypeContext {
	var p = new(VariableSizedTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_variableSizedType
	return p
}

func (*VariableSizedTypeContext) IsVariableSizedTypeContext() {}

func NewVariableSizedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableSizedTypeContext {
	var p = new(VariableSizedTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_variableSizedType

	return p
}

func (s *VariableSizedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableSizedTypeContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *VariableSizedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableSizedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableSizedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterVariableSizedType(s)
	}
}

func (s *VariableSizedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitVariableSizedType(s)
	}
}

func (s *VariableSizedTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitVariableSizedType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) VariableSizedType() (localctx IVariableSizedTypeContext) {
	localctx = NewVariableSizedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, StrictusParserRULE_variableSizedType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.Match(StrictusParserT__5)
	}
	{
		p.SetState(362)
		p.FullType()
	}
	{
		p.SetState(363)
		p.Match(StrictusParserT__6)
	}

	return localctx
}

// IConstantSizedTypeContext is an interface to support dynamic dispatch.
type IConstantSizedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSize returns the size token.
	GetSize() antlr.Token

	// SetSize sets the size token.
	SetSize(antlr.Token)

	// IsConstantSizedTypeContext differentiates from other interfaces.
	IsConstantSizedTypeContext()
}

type ConstantSizedTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	size   antlr.Token
}

func NewEmptyConstantSizedTypeContext() *ConstantSizedTypeContext {
	var p = new(ConstantSizedTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_constantSizedType
	return p
}

func (*ConstantSizedTypeContext) IsConstantSizedTypeContext() {}

func NewConstantSizedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantSizedTypeContext {
	var p = new(ConstantSizedTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_constantSizedType

	return p
}

func (s *ConstantSizedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantSizedTypeContext) GetSize() antlr.Token { return s.size }

func (s *ConstantSizedTypeContext) SetSize(v antlr.Token) { s.size = v }

func (s *ConstantSizedTypeContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *ConstantSizedTypeContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserDecimalLiteral, 0)
}

func (s *ConstantSizedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantSizedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantSizedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterConstantSizedType(s)
	}
}

func (s *ConstantSizedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitConstantSizedType(s)
	}
}

func (s *ConstantSizedTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitConstantSizedType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ConstantSizedType() (localctx IConstantSizedTypeContext) {
	localctx = NewConstantSizedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, StrictusParserRULE_constantSizedType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(StrictusParserT__5)
	}
	{
		p.SetState(366)
		p.FullType()
	}
	{
		p.SetState(367)
		p.Match(StrictusParserT__0)
	}
	{
		p.SetState(368)

		var _m = p.Match(StrictusParserDecimalLiteral)

		localctx.(*ConstantSizedTypeContext).size = _m
	}
	{
		p.SetState(369)
		p.Match(StrictusParserT__6)
	}

	return localctx
}

// IDictionaryTypeContext is an interface to support dynamic dispatch.
type IDictionaryTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKeyType returns the keyType rule contexts.
	GetKeyType() IFullTypeContext

	// GetValueType returns the valueType rule contexts.
	GetValueType() IFullTypeContext

	// SetKeyType sets the keyType rule contexts.
	SetKeyType(IFullTypeContext)

	// SetValueType sets the valueType rule contexts.
	SetValueType(IFullTypeContext)

	// IsDictionaryTypeContext differentiates from other interfaces.
	IsDictionaryTypeContext()
}

type DictionaryTypeContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	keyType   IFullTypeContext
	valueType IFullTypeContext
}

func NewEmptyDictionaryTypeContext() *DictionaryTypeContext {
	var p = new(DictionaryTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_dictionaryType
	return p
}

func (*DictionaryTypeContext) IsDictionaryTypeContext() {}

func NewDictionaryTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictionaryTypeContext {
	var p = new(DictionaryTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_dictionaryType

	return p
}

func (s *DictionaryTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DictionaryTypeContext) GetKeyType() IFullTypeContext { return s.keyType }

func (s *DictionaryTypeContext) GetValueType() IFullTypeContext { return s.valueType }

func (s *DictionaryTypeContext) SetKeyType(v IFullTypeContext) { s.keyType = v }

func (s *DictionaryTypeContext) SetValueType(v IFullTypeContext) { s.valueType = v }

func (s *DictionaryTypeContext) AllFullType() []IFullTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFullTypeContext)(nil)).Elem())
	var tst = make([]IFullTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFullTypeContext)
		}
	}

	return tst
}

func (s *DictionaryTypeContext) FullType(i int) IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *DictionaryTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictionaryTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictionaryTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterDictionaryType(s)
	}
}

func (s *DictionaryTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitDictionaryType(s)
	}
}

func (s *DictionaryTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitDictionaryType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) DictionaryType() (localctx IDictionaryTypeContext) {
	localctx = NewDictionaryTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, StrictusParserRULE_dictionaryType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Match(StrictusParserT__2)
	}
	{
		p.SetState(372)

		var _x = p.FullType()

		localctx.(*DictionaryTypeContext).keyType = _x
	}
	{
		p.SetState(373)
		p.Match(StrictusParserT__4)
	}
	{
		p.SetState(374)

		var _x = p.FullType()

		localctx.(*DictionaryTypeContext).valueType = _x
	}
	{
		p.SetState(375)
		p.Match(StrictusParserT__3)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, StrictusParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.Match(StrictusParserT__2)
	}
	{
		p.SetState(378)
		p.Statements()
	}
	{
		p.SetState(379)
		p.Match(StrictusParserT__3)
	}

	return localctx
}

// IFunctionBlockContext is an interface to support dynamic dispatch.
type IFunctionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBlockContext differentiates from other interfaces.
	IsFunctionBlockContext()
}

type FunctionBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBlockContext() *FunctionBlockContext {
	var p = new(FunctionBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_functionBlock
	return p
}

func (*FunctionBlockContext) IsFunctionBlockContext() {}

func NewFunctionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBlockContext {
	var p = new(FunctionBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_functionBlock

	return p
}

func (s *FunctionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBlockContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *FunctionBlockContext) PreConditions() IPreConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPreConditionsContext)
}

func (s *FunctionBlockContext) PostConditions() IPostConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostConditionsContext)
}

func (s *FunctionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFunctionBlock(s)
	}
}

func (s *FunctionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFunctionBlock(s)
	}
}

func (s *FunctionBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFunctionBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) FunctionBlock() (localctx IFunctionBlockContext) {
	localctx = NewFunctionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, StrictusParserRULE_functionBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match(StrictusParserT__2)
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserPre {
		{
			p.SetState(382)
			p.PreConditions()
		}

	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserPost {
		{
			p.SetState(385)
			p.PostConditions()
		}

	}
	{
		p.SetState(388)
		p.Statements()
	}
	{
		p.SetState(389)
		p.Match(StrictusParserT__3)
	}

	return localctx
}

// IPreConditionsContext is an interface to support dynamic dispatch.
type IPreConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPreConditionsContext differentiates from other interfaces.
	IsPreConditionsContext()
}

type PreConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreConditionsContext() *PreConditionsContext {
	var p = new(PreConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_preConditions
	return p
}

func (*PreConditionsContext) IsPreConditionsContext() {}

func NewPreConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreConditionsContext {
	var p = new(PreConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_preConditions

	return p
}

func (s *PreConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *PreConditionsContext) Pre() antlr.TerminalNode {
	return s.GetToken(StrictusParserPre, 0)
}

func (s *PreConditionsContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *PreConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterPreConditions(s)
	}
}

func (s *PreConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitPreConditions(s)
	}
}

func (s *PreConditionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitPreConditions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PreConditions() (localctx IPreConditionsContext) {
	localctx = NewPreConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, StrictusParserRULE_preConditions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(StrictusParserPre)
	}
	{
		p.SetState(392)
		p.Match(StrictusParserT__2)
	}
	{
		p.SetState(393)
		p.Conditions()
	}
	{
		p.SetState(394)
		p.Match(StrictusParserT__3)
	}

	return localctx
}

// IPostConditionsContext is an interface to support dynamic dispatch.
type IPostConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostConditionsContext differentiates from other interfaces.
	IsPostConditionsContext()
}

type PostConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostConditionsContext() *PostConditionsContext {
	var p = new(PostConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_postConditions
	return p
}

func (*PostConditionsContext) IsPostConditionsContext() {}

func NewPostConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostConditionsContext {
	var p = new(PostConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_postConditions

	return p
}

func (s *PostConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *PostConditionsContext) Post() antlr.TerminalNode {
	return s.GetToken(StrictusParserPost, 0)
}

func (s *PostConditionsContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *PostConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterPostConditions(s)
	}
}

func (s *PostConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitPostConditions(s)
	}
}

func (s *PostConditionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitPostConditions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PostConditions() (localctx IPostConditionsContext) {
	localctx = NewPostConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, StrictusParserRULE_postConditions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(396)
		p.Match(StrictusParserPost)
	}
	{
		p.SetState(397)
		p.Match(StrictusParserT__2)
	}
	{
		p.SetState(398)
		p.Conditions()
	}
	{
		p.SetState(399)
		p.Match(StrictusParserT__3)
	}

	return localctx
}

// IConditionsContext is an interface to support dynamic dispatch.
type IConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionsContext differentiates from other interfaces.
	IsConditionsContext()
}

type ConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsContext() *ConditionsContext {
	var p = new(ConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_conditions
	return p
}

func (*ConditionsContext) IsConditionsContext() {}

func NewConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsContext {
	var p = new(ConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_conditions

	return p
}

func (s *ConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsContext) AllCondition() []IConditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionContext)(nil)).Elem())
	var tst = make([]IConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionContext)
		}
	}

	return tst
}

func (s *ConditionsContext) Condition(i int) IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionsContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *ConditionsContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterConditions(s)
	}
}

func (s *ConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitConditions(s)
	}
}

func (s *ConditionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitConditions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Conditions() (localctx IConditionsContext) {
	localctx = NewConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, StrictusParserRULE_conditions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__2)|(1<<StrictusParserT__5)|(1<<StrictusParserMinus)|(1<<StrictusParserAmpersand)|(1<<StrictusParserNegate)|(1<<StrictusParserMove)|(1<<StrictusParserOpenParen))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(StrictusParserFun-38))|(1<<(StrictusParserTrue-38))|(1<<(StrictusParserFalse-38))|(1<<(StrictusParserNil-38))|(1<<(StrictusParserFrom-38))|(1<<(StrictusParserCreate-38))|(1<<(StrictusParserDestroy-38))|(1<<(StrictusParserIdentifier-38))|(1<<(StrictusParserDecimalLiteral-38))|(1<<(StrictusParserBinaryLiteral-38))|(1<<(StrictusParserOctalLiteral-38))|(1<<(StrictusParserHexadecimalLiteral-38))|(1<<(StrictusParserInvalidNumberLiteral-38))|(1<<(StrictusParserStringLiteral-38)))) != 0) {
		{
			p.SetState(401)
			p.Condition()
		}
		{
			p.SetState(402)
			p.Eos()
		}

		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTest returns the test rule contexts.
	GetTest() IExpressionContext

	// GetMessage returns the message rule contexts.
	GetMessage() IExpressionContext

	// SetTest sets the test rule contexts.
	SetTest(IExpressionContext)

	// SetMessage sets the message rule contexts.
	SetMessage(IExpressionContext)

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	test    IExpressionContext
	message IExpressionContext
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) GetTest() IExpressionContext { return s.test }

func (s *ConditionContext) GetMessage() IExpressionContext { return s.message }

func (s *ConditionContext) SetTest(v IExpressionContext) { s.test = v }

func (s *ConditionContext) SetMessage(v IExpressionContext) { s.message = v }

func (s *ConditionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, StrictusParserRULE_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)

		var _x = p.Expression()

		localctx.(*ConditionContext).test = _x
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(410)
			p.Match(StrictusParserT__4)
		}
		{
			p.SetState(411)

			var _x = p.Expression()

			localctx.(*ConditionContext).message = _x
		}

	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *StatementsContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, StrictusParserRULE_statements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-3)&-(0x1f+1)) == 0 && ((1<<uint((_la-3)))&((1<<(StrictusParserT__2-3))|(1<<(StrictusParserT__5-3))|(1<<(StrictusParserMinus-3))|(1<<(StrictusParserAmpersand-3))|(1<<(StrictusParserNegate-3))|(1<<(StrictusParserMove-3))|(1<<(StrictusParserOpenParen-3))|(1<<(StrictusParserStruct-3)))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(StrictusParserResource-35))|(1<<(StrictusParserContract-35))|(1<<(StrictusParserFun-35))|(1<<(StrictusParserEvent-35))|(1<<(StrictusParserEmit-35))|(1<<(StrictusParserPub-35))|(1<<(StrictusParserPubSet-35))|(1<<(StrictusParserReturn-35))|(1<<(StrictusParserBreak-35))|(1<<(StrictusParserContinue-35))|(1<<(StrictusParserLet-35))|(1<<(StrictusParserVar-35))|(1<<(StrictusParserIf-35))|(1<<(StrictusParserWhile-35))|(1<<(StrictusParserTrue-35))|(1<<(StrictusParserFalse-35))|(1<<(StrictusParserNil-35))|(1<<(StrictusParserImport-35))|(1<<(StrictusParserFrom-35))|(1<<(StrictusParserCreate-35))|(1<<(StrictusParserDestroy-35))|(1<<(StrictusParserIdentifier-35))|(1<<(StrictusParserDecimalLiteral-35))|(1<<(StrictusParserBinaryLiteral-35))|(1<<(StrictusParserOctalLiteral-35))|(1<<(StrictusParserHexadecimalLiteral-35))|(1<<(StrictusParserInvalidNumberLiteral-35))|(1<<(StrictusParserStringLiteral-35)))) != 0) {
		{
			p.SetState(414)
			p.Statement()
		}
		{
			p.SetState(415)
			p.Eos()
		}

		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinueStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) EmitStatement() IEmitStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmitStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmitStatementContext)
}

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Swap() ISwapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwapContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, StrictusParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(422)
			p.ReturnStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(423)
			p.BreakStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(424)
			p.ContinueStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(425)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(426)
			p.WhileStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(427)
			p.EmitStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(428)
			p.Declaration()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(429)
			p.Assignment()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(430)
			p.Swap()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(431)
			p.Expression()
		}

	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(StrictusParserReturn, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, StrictusParserRULE_returnStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		p.Match(StrictusParserReturn)
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		p.SetState(435)

		if !(!p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "!p.lineTerminatorAhead()", ""))
		}
		{
			p.SetState(436)
			p.Expression()
		}

	}

	return localctx
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_breakStatement
	return p
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStatementContext) Break() antlr.TerminalNode {
	return s.GetToken(StrictusParserBreak, 0)
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, StrictusParserRULE_breakStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)
		p.Match(StrictusParserBreak)
	}

	return localctx
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_continueStatement
	return p
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(StrictusParserContinue, 0)
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, StrictusParserRULE_continueStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(441)
		p.Match(StrictusParserContinue)
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTestExpression returns the testExpression rule contexts.
	GetTestExpression() IExpressionContext

	// GetTestDeclaration returns the testDeclaration rule contexts.
	GetTestDeclaration() IVariableDeclarationContext

	// GetThen returns the then rule contexts.
	GetThen() IBlockContext

	// GetAlt returns the alt rule contexts.
	GetAlt() IBlockContext

	// SetTestExpression sets the testExpression rule contexts.
	SetTestExpression(IExpressionContext)

	// SetTestDeclaration sets the testDeclaration rule contexts.
	SetTestDeclaration(IVariableDeclarationContext)

	// SetThen sets the then rule contexts.
	SetThen(IBlockContext)

	// SetAlt sets the alt rule contexts.
	SetAlt(IBlockContext)

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	testExpression  IExpressionContext
	testDeclaration IVariableDeclarationContext
	then            IBlockContext
	alt             IBlockContext
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) GetTestExpression() IExpressionContext { return s.testExpression }

func (s *IfStatementContext) GetTestDeclaration() IVariableDeclarationContext {
	return s.testDeclaration
}

func (s *IfStatementContext) GetThen() IBlockContext { return s.then }

func (s *IfStatementContext) GetAlt() IBlockContext { return s.alt }

func (s *IfStatementContext) SetTestExpression(v IExpressionContext) { s.testExpression = v }

func (s *IfStatementContext) SetTestDeclaration(v IVariableDeclarationContext) { s.testDeclaration = v }

func (s *IfStatementContext) SetThen(v IBlockContext) { s.then = v }

func (s *IfStatementContext) SetAlt(v IBlockContext) { s.alt = v }

func (s *IfStatementContext) If() antlr.TerminalNode {
	return s.GetToken(StrictusParserIf, 0)
}

func (s *IfStatementContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *IfStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(StrictusParserElse, 0)
}

func (s *IfStatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, StrictusParserRULE_ifStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(443)
		p.Match(StrictusParserIf)
	}
	p.SetState(446)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserT__2, StrictusParserT__5, StrictusParserMinus, StrictusParserAmpersand, StrictusParserNegate, StrictusParserMove, StrictusParserOpenParen, StrictusParserFun, StrictusParserTrue, StrictusParserFalse, StrictusParserNil, StrictusParserFrom, StrictusParserCreate, StrictusParserDestroy, StrictusParserIdentifier, StrictusParserDecimalLiteral, StrictusParserBinaryLiteral, StrictusParserOctalLiteral, StrictusParserHexadecimalLiteral, StrictusParserInvalidNumberLiteral, StrictusParserStringLiteral:
		{
			p.SetState(444)

			var _x = p.Expression()

			localctx.(*IfStatementContext).testExpression = _x
		}

	case StrictusParserLet, StrictusParserVar:
		{
			p.SetState(445)

			var _x = p.VariableDeclaration()

			localctx.(*IfStatementContext).testDeclaration = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(448)

		var _x = p.Block()

		localctx.(*IfStatementContext).then = _x
	}
	p.SetState(454)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(449)
			p.Match(StrictusParserElse)
		}
		p.SetState(452)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case StrictusParserIf:
			{
				p.SetState(450)
				p.IfStatement()
			}

		case StrictusParserT__2:
			{
				p.SetState(451)

				var _x = p.Block()

				localctx.(*IfStatementContext).alt = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_whileStatement
	return p
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(StrictusParserWhile, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, StrictusParserRULE_whileStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(456)
		p.Match(StrictusParserWhile)
	}
	{
		p.SetState(457)
		p.Expression()
	}
	{
		p.SetState(458)
		p.Block()
	}

	return localctx
}

// IEmitStatementContext is an interface to support dynamic dispatch.
type IEmitStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmitStatementContext differentiates from other interfaces.
	IsEmitStatementContext()
}

type EmitStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmitStatementContext() *EmitStatementContext {
	var p = new(EmitStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_emitStatement
	return p
}

func (*EmitStatementContext) IsEmitStatementContext() {}

func NewEmitStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmitStatementContext {
	var p = new(EmitStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_emitStatement

	return p
}

func (s *EmitStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *EmitStatementContext) Emit() antlr.TerminalNode {
	return s.GetToken(StrictusParserEmit, 0)
}

func (s *EmitStatementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *EmitStatementContext) Invocation() IInvocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *EmitStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmitStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmitStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterEmitStatement(s)
	}
}

func (s *EmitStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitEmitStatement(s)
	}
}

func (s *EmitStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitEmitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) EmitStatement() (localctx IEmitStatementContext) {
	localctx = NewEmitStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, StrictusParserRULE_emitStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		p.Match(StrictusParserEmit)
	}
	{
		p.SetState(461)
		p.Identifier()
	}
	{
		p.SetState(462)
		p.Invocation()
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) VariableKind() IVariableKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableKindContext)
}

func (s *VariableDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *VariableDeclarationContext) Transfer() ITransferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransferContext)
}

func (s *VariableDeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, StrictusParserRULE_variableDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		p.VariableKind()
	}
	{
		p.SetState(465)
		p.Identifier()
	}
	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserT__4 {
		{
			p.SetState(466)
			p.Match(StrictusParserT__4)
		}
		{
			p.SetState(467)
			p.TypeAnnotation()
		}

	}
	{
		p.SetState(470)
		p.Transfer()
	}
	{
		p.SetState(471)
		p.Expression()
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AssignmentContext) Transfer() ITransferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransferContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) AllExpressionAccess() []IExpressionAccessContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAccessContext)(nil)).Elem())
	var tst = make([]IExpressionAccessContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAccessContext)
		}
	}

	return tst
}

func (s *AssignmentContext) ExpressionAccess(i int) IExpressionAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAccessContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAccessContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, StrictusParserRULE_assignment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(473)
		p.Identifier()
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == StrictusParserT__5 || _la == StrictusParserT__11 {
		{
			p.SetState(474)
			p.ExpressionAccess()
		}

		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(480)
		p.Transfer()
	}
	{
		p.SetState(481)
		p.Expression()
	}

	return localctx
}

// ISwapContext is an interface to support dynamic dispatch.
type ISwapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionContext)

	// IsSwapContext differentiates from other interfaces.
	IsSwapContext()
}

type SwapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpressionContext
	right  IExpressionContext
}

func NewEmptySwapContext() *SwapContext {
	var p = new(SwapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_swap
	return p
}

func (*SwapContext) IsSwapContext() {}

func NewSwapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwapContext {
	var p = new(SwapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_swap

	return p
}

func (s *SwapContext) GetParser() antlr.Parser { return s.parser }

func (s *SwapContext) GetLeft() IExpressionContext { return s.left }

func (s *SwapContext) GetRight() IExpressionContext { return s.right }

func (s *SwapContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *SwapContext) SetRight(v IExpressionContext) { s.right = v }

func (s *SwapContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *SwapContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterSwap(s)
	}
}

func (s *SwapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitSwap(s)
	}
}

func (s *SwapContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitSwap(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Swap() (localctx ISwapContext) {
	localctx = NewSwapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, StrictusParserRULE_swap)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)

		var _x = p.Expression()

		localctx.(*SwapContext).left = _x
	}
	{
		p.SetState(484)
		p.Match(StrictusParserT__7)
	}
	{
		p.SetState(485)

		var _x = p.Expression()

		localctx.(*SwapContext).right = _x
	}

	return localctx
}

// ITransferContext is an interface to support dynamic dispatch.
type ITransferContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransferContext differentiates from other interfaces.
	IsTransferContext()
}

type TransferContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransferContext() *TransferContext {
	var p = new(TransferContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_transfer
	return p
}

func (*TransferContext) IsTransferContext() {}

func NewTransferContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransferContext {
	var p = new(TransferContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_transfer

	return p
}

func (s *TransferContext) GetParser() antlr.Parser { return s.parser }

func (s *TransferContext) Move() antlr.TerminalNode {
	return s.GetToken(StrictusParserMove, 0)
}

func (s *TransferContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransferContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterTransfer(s)
	}
}

func (s *TransferContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitTransfer(s)
	}
}

func (s *TransferContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitTransfer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Transfer() (localctx ITransferContext) {
	localctx = NewTransferContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, StrictusParserRULE_transfer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(487)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserT__8 || _la == StrictusParserMove) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ConditionalExpression() IConditionalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, StrictusParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(489)
		p.ConditionalExpression()
	}

	return localctx
}

// IConditionalExpressionContext is an interface to support dynamic dispatch.
type IConditionalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetThen returns the then rule contexts.
	GetThen() IExpressionContext

	// GetAlt returns the alt rule contexts.
	GetAlt() IExpressionContext

	// SetThen sets the then rule contexts.
	SetThen(IExpressionContext)

	// SetAlt sets the alt rule contexts.
	SetAlt(IExpressionContext)

	// IsConditionalExpressionContext differentiates from other interfaces.
	IsConditionalExpressionContext()
}

type ConditionalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	then   IExpressionContext
	alt    IExpressionContext
}

func NewEmptyConditionalExpressionContext() *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_conditionalExpression
	return p
}

func (*ConditionalExpressionContext) IsConditionalExpressionContext() {}

func NewConditionalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_conditionalExpression

	return p
}

func (s *ConditionalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalExpressionContext) GetThen() IExpressionContext { return s.then }

func (s *ConditionalExpressionContext) GetAlt() IExpressionContext { return s.alt }

func (s *ConditionalExpressionContext) SetThen(v IExpressionContext) { s.then = v }

func (s *ConditionalExpressionContext) SetAlt(v IExpressionContext) { s.alt = v }

func (s *ConditionalExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *ConditionalExpressionContext) Optional() antlr.TerminalNode {
	return s.GetToken(StrictusParserOptional, 0)
}

func (s *ConditionalExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionalExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitConditionalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ConditionalExpression() (localctx IConditionalExpressionContext) {
	localctx = NewConditionalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, StrictusParserRULE_conditionalExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
		p.orExpression(0)
	}
	p.SetState(497)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(492)
			p.Match(StrictusParserOptional)
		}
		{
			p.SetState(493)

			var _x = p.Expression()

			localctx.(*ConditionalExpressionContext).then = _x
		}
		{
			p.SetState(494)
			p.Match(StrictusParserT__4)
		}
		{
			p.SetState(495)

			var _x = p.Expression()

			localctx.(*ConditionalExpressionContext).alt = _x
		}

	}

	return localctx
}

// IOrExpressionContext is an interface to support dynamic dispatch.
type IOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrExpressionContext differentiates from other interfaces.
	IsOrExpressionContext()
}

type OrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExpressionContext() *OrExpressionContext {
	var p = new(OrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_orExpression
	return p
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_orExpression

	return p
}

func (s *OrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *OrExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (s *OrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) OrExpression() (localctx IOrExpressionContext) {
	return p.orExpression(0)
}

func (p *StrictusParser) orExpression(_p int) (localctx IOrExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IOrExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 92
	p.EnterRecursionRule(localctx, 92, StrictusParserRULE_orExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)
		p.andExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOrExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_orExpression)
			p.SetState(502)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(503)
				p.Match(StrictusParserT__9)
			}
			{
				p.SetState(504)
				p.andExpression(0)
			}

		}
		p.SetState(509)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
	}

	return localctx
}

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_andExpression
	return p
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *AndExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) AndExpression() (localctx IAndExpressionContext) {
	return p.andExpression(0)
}

func (p *StrictusParser) andExpression(_p int) (localctx IAndExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAndExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 94
	p.EnterRecursionRule(localctx, 94, StrictusParserRULE_andExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(511)
		p.equalityExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_andExpression)
			p.SetState(513)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(514)
				p.Match(StrictusParserT__10)
			}
			{
				p.SetState(515)
				p.equalityExpression(0)
			}

		}
		p.SetState(520)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}

	return localctx
}

// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_equalityExpression
	return p
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *EqualityExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *EqualityExpressionContext) EqualityOp() IEqualityOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityOpContext)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	return p.equalityExpression(0)
}

func (p *StrictusParser) equalityExpression(_p int) (localctx IEqualityExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEqualityExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 96
	p.EnterRecursionRule(localctx, 96, StrictusParserRULE_equalityExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(522)
		p.relationalExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(530)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewEqualityExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_equalityExpression)
			p.SetState(524)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(525)
				p.EqualityOp()
			}
			{
				p.SetState(526)
				p.relationalExpression(0)
			}

		}
		p.SetState(532)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
	}

	return localctx
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_relationalExpression
	return p
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) NilCoalescingExpression() INilCoalescingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INilCoalescingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INilCoalescingExpressionContext)
}

func (s *RelationalExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *RelationalExpressionContext) RelationalOp() IRelationalOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalOpContext)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	return p.relationalExpression(0)
}

func (p *StrictusParser) relationalExpression(_p int) (localctx IRelationalExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelationalExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 98
	p.EnterRecursionRule(localctx, 98, StrictusParserRULE_relationalExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.NilCoalescingExpression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(542)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_relationalExpression)
			p.SetState(536)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(537)
				p.RelationalOp()
			}
			{
				p.SetState(538)
				p.NilCoalescingExpression()
			}

		}
		p.SetState(544)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
	}

	return localctx
}

// INilCoalescingExpressionContext is an interface to support dynamic dispatch.
type INilCoalescingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNilCoalescingExpressionContext differentiates from other interfaces.
	IsNilCoalescingExpressionContext()
}

type NilCoalescingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNilCoalescingExpressionContext() *NilCoalescingExpressionContext {
	var p = new(NilCoalescingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_nilCoalescingExpression
	return p
}

func (*NilCoalescingExpressionContext) IsNilCoalescingExpressionContext() {}

func NewNilCoalescingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NilCoalescingExpressionContext {
	var p = new(NilCoalescingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_nilCoalescingExpression

	return p
}

func (s *NilCoalescingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NilCoalescingExpressionContext) FailableDowncastingExpression() IFailableDowncastingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFailableDowncastingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFailableDowncastingExpressionContext)
}

func (s *NilCoalescingExpressionContext) NilCoalescing() antlr.TerminalNode {
	return s.GetToken(StrictusParserNilCoalescing, 0)
}

func (s *NilCoalescingExpressionContext) NilCoalescingExpression() INilCoalescingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INilCoalescingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INilCoalescingExpressionContext)
}

func (s *NilCoalescingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilCoalescingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NilCoalescingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterNilCoalescingExpression(s)
	}
}

func (s *NilCoalescingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitNilCoalescingExpression(s)
	}
}

func (s *NilCoalescingExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitNilCoalescingExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) NilCoalescingExpression() (localctx INilCoalescingExpressionContext) {
	localctx = NewNilCoalescingExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, StrictusParserRULE_nilCoalescingExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(545)
		p.failableDowncastingExpression(0)
	}
	p.SetState(548)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(546)
			p.Match(StrictusParserNilCoalescing)
		}
		{
			p.SetState(547)
			p.NilCoalescingExpression()
		}

	}

	return localctx
}

// IFailableDowncastingExpressionContext is an interface to support dynamic dispatch.
type IFailableDowncastingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFailableDowncastingExpressionContext differentiates from other interfaces.
	IsFailableDowncastingExpressionContext()
}

type FailableDowncastingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFailableDowncastingExpressionContext() *FailableDowncastingExpressionContext {
	var p = new(FailableDowncastingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_failableDowncastingExpression
	return p
}

func (*FailableDowncastingExpressionContext) IsFailableDowncastingExpressionContext() {}

func NewFailableDowncastingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FailableDowncastingExpressionContext {
	var p = new(FailableDowncastingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_failableDowncastingExpression

	return p
}

func (s *FailableDowncastingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FailableDowncastingExpressionContext) ConcatenatingExpression() IConcatenatingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcatenatingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConcatenatingExpressionContext)
}

func (s *FailableDowncastingExpressionContext) FailableDowncastingExpression() IFailableDowncastingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFailableDowncastingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFailableDowncastingExpressionContext)
}

func (s *FailableDowncastingExpressionContext) FailableDowncasting() antlr.TerminalNode {
	return s.GetToken(StrictusParserFailableDowncasting, 0)
}

func (s *FailableDowncastingExpressionContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *FailableDowncastingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FailableDowncastingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FailableDowncastingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFailableDowncastingExpression(s)
	}
}

func (s *FailableDowncastingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFailableDowncastingExpression(s)
	}
}

func (s *FailableDowncastingExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFailableDowncastingExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) FailableDowncastingExpression() (localctx IFailableDowncastingExpressionContext) {
	return p.failableDowncastingExpression(0)
}

func (p *StrictusParser) failableDowncastingExpression(_p int) (localctx IFailableDowncastingExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFailableDowncastingExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFailableDowncastingExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 102
	p.EnterRecursionRule(localctx, 102, StrictusParserRULE_failableDowncastingExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(551)
		p.concatenatingExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFailableDowncastingExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_failableDowncastingExpression)
			p.SetState(553)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(554)
				p.Match(StrictusParserFailableDowncasting)
			}
			{
				p.SetState(555)
				p.TypeAnnotation()
			}

		}
		p.SetState(560)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}

	return localctx
}

// IConcatenatingExpressionContext is an interface to support dynamic dispatch.
type IConcatenatingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConcatenatingExpressionContext differentiates from other interfaces.
	IsConcatenatingExpressionContext()
}

type ConcatenatingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcatenatingExpressionContext() *ConcatenatingExpressionContext {
	var p = new(ConcatenatingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_concatenatingExpression
	return p
}

func (*ConcatenatingExpressionContext) IsConcatenatingExpressionContext() {}

func NewConcatenatingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcatenatingExpressionContext {
	var p = new(ConcatenatingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_concatenatingExpression

	return p
}

func (s *ConcatenatingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcatenatingExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *ConcatenatingExpressionContext) ConcatenatingExpression() IConcatenatingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcatenatingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConcatenatingExpressionContext)
}

func (s *ConcatenatingExpressionContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(StrictusParserAmpersand, 0)
}

func (s *ConcatenatingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatenatingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConcatenatingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterConcatenatingExpression(s)
	}
}

func (s *ConcatenatingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitConcatenatingExpression(s)
	}
}

func (s *ConcatenatingExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitConcatenatingExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ConcatenatingExpression() (localctx IConcatenatingExpressionContext) {
	return p.concatenatingExpression(0)
}

func (p *StrictusParser) concatenatingExpression(_p int) (localctx IConcatenatingExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewConcatenatingExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConcatenatingExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 104
	p.EnterRecursionRule(localctx, 104, StrictusParserRULE_concatenatingExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(562)
		p.additiveExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(569)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewConcatenatingExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_concatenatingExpression)
			p.SetState(564)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(565)
				p.Match(StrictusParserAmpersand)
			}
			{
				p.SetState(566)
				p.additiveExpression(0)
			}

		}
		p.SetState(571)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
	}

	return localctx
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_additiveExpression
	return p
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) MultiplicativeExpression() IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *AdditiveExpressionContext) AdditiveOp() IAdditiveOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveOpContext)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	return p.additiveExpression(0)
}

func (p *StrictusParser) additiveExpression(_p int) (localctx IAdditiveExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAdditiveExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 106
	p.EnterRecursionRule(localctx, 106, StrictusParserRULE_additiveExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(573)
		p.multiplicativeExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(581)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAdditiveExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_additiveExpression)
			p.SetState(575)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(576)
				p.AdditiveOp()
			}
			{
				p.SetState(577)
				p.multiplicativeExpression(0)
			}

		}
		p.SetState(583)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())
	}

	return localctx
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_multiplicativeExpression
	return p
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicativeExpressionContext) MultiplicativeExpression() IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *MultiplicativeExpressionContext) MultiplicativeOp() IMultiplicativeOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeOpContext)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	return p.multiplicativeExpression(0)
}

func (p *StrictusParser) multiplicativeExpression(_p int) (localctx IMultiplicativeExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultiplicativeExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 108
	p.EnterRecursionRule(localctx, 108, StrictusParserRULE_multiplicativeExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(585)
		p.UnaryExpression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(593)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMultiplicativeExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_multiplicativeExpression)
			p.SetState(587)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(588)
				p.MultiplicativeOp()
			}
			{
				p.SetState(589)
				p.UnaryExpression()
			}

		}
		p.SetState(595)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_unaryExpression
	return p
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *UnaryExpressionContext) AllUnaryOp() []IUnaryOpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnaryOpContext)(nil)).Elem())
	var tst = make([]IUnaryOpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnaryOpContext)
		}
	}

	return tst
}

func (s *UnaryExpressionContext) UnaryOp(i int) IUnaryOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, StrictusParserRULE_unaryExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(604)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(596)
			p.PrimaryExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(598)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(597)
					p.UnaryOp()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(600)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())
		}
		{
			p.SetState(602)
			p.UnaryExpression()
		}

	}

	return localctx
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_primaryExpression
	return p
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) CreateExpression() ICreateExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreateExpressionContext)
}

func (s *PrimaryExpressionContext) DestroyExpression() IDestroyExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestroyExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDestroyExpressionContext)
}

func (s *PrimaryExpressionContext) ReferenceExpression() IReferenceExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferenceExpressionContext)
}

func (s *PrimaryExpressionContext) ComposedExpression() IComposedExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComposedExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComposedExpressionContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, StrictusParserRULE_primaryExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(606)
			p.CreateExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(607)
			p.DestroyExpression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(608)
			p.ReferenceExpression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(609)
			p.ComposedExpression()
		}

	}

	return localctx
}

// IComposedExpressionContext is an interface to support dynamic dispatch.
type IComposedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComposedExpressionContext differentiates from other interfaces.
	IsComposedExpressionContext()
}

type ComposedExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComposedExpressionContext() *ComposedExpressionContext {
	var p = new(ComposedExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_composedExpression
	return p
}

func (*ComposedExpressionContext) IsComposedExpressionContext() {}

func NewComposedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComposedExpressionContext {
	var p = new(ComposedExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_composedExpression

	return p
}

func (s *ComposedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ComposedExpressionContext) PrimaryExpressionStart() IPrimaryExpressionStartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionStartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionStartContext)
}

func (s *ComposedExpressionContext) AllPrimaryExpressionSuffix() []IPrimaryExpressionSuffixContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimaryExpressionSuffixContext)(nil)).Elem())
	var tst = make([]IPrimaryExpressionSuffixContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimaryExpressionSuffixContext)
		}
	}

	return tst
}

func (s *ComposedExpressionContext) PrimaryExpressionSuffix(i int) IPrimaryExpressionSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionSuffixContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionSuffixContext)
}

func (s *ComposedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComposedExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComposedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterComposedExpression(s)
	}
}

func (s *ComposedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitComposedExpression(s)
	}
}

func (s *ComposedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitComposedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ComposedExpression() (localctx IComposedExpressionContext) {
	localctx = NewComposedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, StrictusParserRULE_composedExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(612)
		p.PrimaryExpressionStart()
	}
	p.SetState(616)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(613)
				p.PrimaryExpressionSuffix()
			}

		}
		p.SetState(618)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExpressionSuffixContext is an interface to support dynamic dispatch.
type IPrimaryExpressionSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionSuffixContext differentiates from other interfaces.
	IsPrimaryExpressionSuffixContext()
}

type PrimaryExpressionSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionSuffixContext() *PrimaryExpressionSuffixContext {
	var p = new(PrimaryExpressionSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_primaryExpressionSuffix
	return p
}

func (*PrimaryExpressionSuffixContext) IsPrimaryExpressionSuffixContext() {}

func NewPrimaryExpressionSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionSuffixContext {
	var p = new(PrimaryExpressionSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_primaryExpressionSuffix

	return p
}

func (s *PrimaryExpressionSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionSuffixContext) ExpressionAccess() IExpressionAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAccessContext)
}

func (s *PrimaryExpressionSuffixContext) Invocation() IInvocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *PrimaryExpressionSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterPrimaryExpressionSuffix(s)
	}
}

func (s *PrimaryExpressionSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitPrimaryExpressionSuffix(s)
	}
}

func (s *PrimaryExpressionSuffixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitPrimaryExpressionSuffix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PrimaryExpressionSuffix() (localctx IPrimaryExpressionSuffixContext) {
	localctx = NewPrimaryExpressionSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, StrictusParserRULE_primaryExpressionSuffix)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(621)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserT__5, StrictusParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(619)
			p.ExpressionAccess()
		}

	case StrictusParserOpenParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(620)
			p.Invocation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEqualityOpContext is an interface to support dynamic dispatch.
type IEqualityOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualityOpContext differentiates from other interfaces.
	IsEqualityOpContext()
}

type EqualityOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityOpContext() *EqualityOpContext {
	var p = new(EqualityOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_equalityOp
	return p
}

func (*EqualityOpContext) IsEqualityOpContext() {}

func NewEqualityOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityOpContext {
	var p = new(EqualityOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_equalityOp

	return p
}

func (s *EqualityOpContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityOpContext) Equal() antlr.TerminalNode {
	return s.GetToken(StrictusParserEqual, 0)
}

func (s *EqualityOpContext) Unequal() antlr.TerminalNode {
	return s.GetToken(StrictusParserUnequal, 0)
}

func (s *EqualityOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterEqualityOp(s)
	}
}

func (s *EqualityOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitEqualityOp(s)
	}
}

func (s *EqualityOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitEqualityOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) EqualityOp() (localctx IEqualityOpContext) {
	localctx = NewEqualityOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, StrictusParserRULE_equalityOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(623)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserEqual || _la == StrictusParserUnequal) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRelationalOpContext is an interface to support dynamic dispatch.
type IRelationalOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalOpContext differentiates from other interfaces.
	IsRelationalOpContext()
}

type RelationalOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalOpContext() *RelationalOpContext {
	var p = new(RelationalOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_relationalOp
	return p
}

func (*RelationalOpContext) IsRelationalOpContext() {}

func NewRelationalOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalOpContext {
	var p = new(RelationalOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_relationalOp

	return p
}

func (s *RelationalOpContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalOpContext) Less() antlr.TerminalNode {
	return s.GetToken(StrictusParserLess, 0)
}

func (s *RelationalOpContext) Greater() antlr.TerminalNode {
	return s.GetToken(StrictusParserGreater, 0)
}

func (s *RelationalOpContext) LessEqual() antlr.TerminalNode {
	return s.GetToken(StrictusParserLessEqual, 0)
}

func (s *RelationalOpContext) GreaterEqual() antlr.TerminalNode {
	return s.GetToken(StrictusParserGreaterEqual, 0)
}

func (s *RelationalOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterRelationalOp(s)
	}
}

func (s *RelationalOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitRelationalOp(s)
	}
}

func (s *RelationalOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitRelationalOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) RelationalOp() (localctx IRelationalOpContext) {
	localctx = NewRelationalOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, StrictusParserRULE_relationalOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(625)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserLess)|(1<<StrictusParserGreater)|(1<<StrictusParserLessEqual)|(1<<StrictusParserGreaterEqual))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAdditiveOpContext is an interface to support dynamic dispatch.
type IAdditiveOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveOpContext differentiates from other interfaces.
	IsAdditiveOpContext()
}

type AdditiveOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveOpContext() *AdditiveOpContext {
	var p = new(AdditiveOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_additiveOp
	return p
}

func (*AdditiveOpContext) IsAdditiveOpContext() {}

func NewAdditiveOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveOpContext {
	var p = new(AdditiveOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_additiveOp

	return p
}

func (s *AdditiveOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveOpContext) Plus() antlr.TerminalNode {
	return s.GetToken(StrictusParserPlus, 0)
}

func (s *AdditiveOpContext) Minus() antlr.TerminalNode {
	return s.GetToken(StrictusParserMinus, 0)
}

func (s *AdditiveOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAdditiveOp(s)
	}
}

func (s *AdditiveOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAdditiveOp(s)
	}
}

func (s *AdditiveOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAdditiveOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) AdditiveOp() (localctx IAdditiveOpContext) {
	localctx = NewAdditiveOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, StrictusParserRULE_additiveOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(627)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserPlus || _la == StrictusParserMinus) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMultiplicativeOpContext is an interface to support dynamic dispatch.
type IMultiplicativeOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeOpContext differentiates from other interfaces.
	IsMultiplicativeOpContext()
}

type MultiplicativeOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeOpContext() *MultiplicativeOpContext {
	var p = new(MultiplicativeOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_multiplicativeOp
	return p
}

func (*MultiplicativeOpContext) IsMultiplicativeOpContext() {}

func NewMultiplicativeOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeOpContext {
	var p = new(MultiplicativeOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_multiplicativeOp

	return p
}

func (s *MultiplicativeOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeOpContext) Mul() antlr.TerminalNode {
	return s.GetToken(StrictusParserMul, 0)
}

func (s *MultiplicativeOpContext) Div() antlr.TerminalNode {
	return s.GetToken(StrictusParserDiv, 0)
}

func (s *MultiplicativeOpContext) Mod() antlr.TerminalNode {
	return s.GetToken(StrictusParserMod, 0)
}

func (s *MultiplicativeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterMultiplicativeOp(s)
	}
}

func (s *MultiplicativeOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitMultiplicativeOp(s)
	}
}

func (s *MultiplicativeOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitMultiplicativeOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) MultiplicativeOp() (localctx IMultiplicativeOpContext) {
	localctx = NewMultiplicativeOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, StrictusParserRULE_multiplicativeOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(629)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserMul)|(1<<StrictusParserDiv)|(1<<StrictusParserMod))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnaryOpContext is an interface to support dynamic dispatch.
type IUnaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOpContext differentiates from other interfaces.
	IsUnaryOpContext()
}

type UnaryOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpContext() *UnaryOpContext {
	var p = new(UnaryOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_unaryOp
	return p
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOpContext) Minus() antlr.TerminalNode {
	return s.GetToken(StrictusParserMinus, 0)
}

func (s *UnaryOpContext) Negate() antlr.TerminalNode {
	return s.GetToken(StrictusParserNegate, 0)
}

func (s *UnaryOpContext) Move() antlr.TerminalNode {
	return s.GetToken(StrictusParserMove, 0)
}

func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterUnaryOp(s)
	}
}

func (s *UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitUnaryOp(s)
	}
}

func (s *UnaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitUnaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) UnaryOp() (localctx IUnaryOpContext) {
	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, StrictusParserRULE_unaryOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(631)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserMinus)|(1<<StrictusParserNegate)|(1<<StrictusParserMove))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPrimaryExpressionStartContext is an interface to support dynamic dispatch.
type IPrimaryExpressionStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionStartContext differentiates from other interfaces.
	IsPrimaryExpressionStartContext()
}

type PrimaryExpressionStartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionStartContext() *PrimaryExpressionStartContext {
	var p = new(PrimaryExpressionStartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_primaryExpressionStart
	return p
}

func (*PrimaryExpressionStartContext) IsPrimaryExpressionStartContext() {}

func NewPrimaryExpressionStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionStartContext {
	var p = new(PrimaryExpressionStartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_primaryExpressionStart

	return p
}

func (s *PrimaryExpressionStartContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionStartContext) IdentifierExpression() IIdentifierExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierExpressionContext)
}

func (s *PrimaryExpressionStartContext) LiteralExpression() ILiteralExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralExpressionContext)
}

func (s *PrimaryExpressionStartContext) FunctionExpression() IFunctionExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionExpressionContext)
}

func (s *PrimaryExpressionStartContext) NestedExpression() INestedExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INestedExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INestedExpressionContext)
}

func (s *PrimaryExpressionStartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionStartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionStartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterPrimaryExpressionStart(s)
	}
}

func (s *PrimaryExpressionStartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitPrimaryExpressionStart(s)
	}
}

func (s *PrimaryExpressionStartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitPrimaryExpressionStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PrimaryExpressionStart() (localctx IPrimaryExpressionStartContext) {
	localctx = NewPrimaryExpressionStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, StrictusParserRULE_primaryExpressionStart)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(637)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserFrom, StrictusParserCreate, StrictusParserDestroy, StrictusParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(633)
			p.IdentifierExpression()
		}

	case StrictusParserT__2, StrictusParserT__5, StrictusParserMinus, StrictusParserTrue, StrictusParserFalse, StrictusParserNil, StrictusParserDecimalLiteral, StrictusParserBinaryLiteral, StrictusParserOctalLiteral, StrictusParserHexadecimalLiteral, StrictusParserInvalidNumberLiteral, StrictusParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(634)
			p.LiteralExpression()
		}

	case StrictusParserFun:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(635)
			p.FunctionExpression()
		}

	case StrictusParserOpenParen:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(636)
			p.NestedExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICreateExpressionContext is an interface to support dynamic dispatch.
type ICreateExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateExpressionContext differentiates from other interfaces.
	IsCreateExpressionContext()
}

type CreateExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateExpressionContext() *CreateExpressionContext {
	var p = new(CreateExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_createExpression
	return p
}

func (*CreateExpressionContext) IsCreateExpressionContext() {}

func NewCreateExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateExpressionContext {
	var p = new(CreateExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_createExpression

	return p
}

func (s *CreateExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateExpressionContext) Create() antlr.TerminalNode {
	return s.GetToken(StrictusParserCreate, 0)
}

func (s *CreateExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *CreateExpressionContext) Invocation() IInvocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *CreateExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterCreateExpression(s)
	}
}

func (s *CreateExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitCreateExpression(s)
	}
}

func (s *CreateExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitCreateExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) CreateExpression() (localctx ICreateExpressionContext) {
	localctx = NewCreateExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, StrictusParserRULE_createExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(639)
		p.Match(StrictusParserCreate)
	}
	{
		p.SetState(640)
		p.Identifier()
	}
	{
		p.SetState(641)
		p.Invocation()
	}

	return localctx
}

// IDestroyExpressionContext is an interface to support dynamic dispatch.
type IDestroyExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDestroyExpressionContext differentiates from other interfaces.
	IsDestroyExpressionContext()
}

type DestroyExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDestroyExpressionContext() *DestroyExpressionContext {
	var p = new(DestroyExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_destroyExpression
	return p
}

func (*DestroyExpressionContext) IsDestroyExpressionContext() {}

func NewDestroyExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestroyExpressionContext {
	var p = new(DestroyExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_destroyExpression

	return p
}

func (s *DestroyExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *DestroyExpressionContext) Destroy() antlr.TerminalNode {
	return s.GetToken(StrictusParserDestroy, 0)
}

func (s *DestroyExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DestroyExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestroyExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestroyExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterDestroyExpression(s)
	}
}

func (s *DestroyExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitDestroyExpression(s)
	}
}

func (s *DestroyExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitDestroyExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) DestroyExpression() (localctx IDestroyExpressionContext) {
	localctx = NewDestroyExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, StrictusParserRULE_destroyExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(643)
		p.Match(StrictusParserDestroy)
	}
	{
		p.SetState(644)
		p.Expression()
	}

	return localctx
}

// IReferenceExpressionContext is an interface to support dynamic dispatch.
type IReferenceExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferenceExpressionContext differentiates from other interfaces.
	IsReferenceExpressionContext()
}

type ReferenceExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceExpressionContext() *ReferenceExpressionContext {
	var p = new(ReferenceExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_referenceExpression
	return p
}

func (*ReferenceExpressionContext) IsReferenceExpressionContext() {}

func NewReferenceExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceExpressionContext {
	var p = new(ReferenceExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_referenceExpression

	return p
}

func (s *ReferenceExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceExpressionContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(StrictusParserAmpersand, 0)
}

func (s *ReferenceExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReferenceExpressionContext) Downcasting() antlr.TerminalNode {
	return s.GetToken(StrictusParserDowncasting, 0)
}

func (s *ReferenceExpressionContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *ReferenceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterReferenceExpression(s)
	}
}

func (s *ReferenceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitReferenceExpression(s)
	}
}

func (s *ReferenceExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitReferenceExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ReferenceExpression() (localctx IReferenceExpressionContext) {
	localctx = NewReferenceExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, StrictusParserRULE_referenceExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(646)
		p.Match(StrictusParserAmpersand)
	}
	{
		p.SetState(647)
		p.Expression()
	}
	{
		p.SetState(648)
		p.Match(StrictusParserDowncasting)
	}
	{
		p.SetState(649)
		p.FullType()
	}

	return localctx
}

// IIdentifierExpressionContext is an interface to support dynamic dispatch.
type IIdentifierExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierExpressionContext differentiates from other interfaces.
	IsIdentifierExpressionContext()
}

type IdentifierExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierExpressionContext() *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_identifierExpression
	return p
}

func (*IdentifierExpressionContext) IsIdentifierExpressionContext() {}

func NewIdentifierExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_identifierExpression

	return p
}

func (s *IdentifierExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) IdentifierExpression() (localctx IIdentifierExpressionContext) {
	localctx = NewIdentifierExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, StrictusParserRULE_identifierExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(651)
		p.Identifier()
	}

	return localctx
}

// ILiteralExpressionContext is an interface to support dynamic dispatch.
type ILiteralExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralExpressionContext differentiates from other interfaces.
	IsLiteralExpressionContext()
}

type LiteralExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralExpressionContext() *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_literalExpression
	return p
}

func (*LiteralExpressionContext) IsLiteralExpressionContext() {}

func NewLiteralExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_literalExpression

	return p
}

func (s *LiteralExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) LiteralExpression() (localctx ILiteralExpressionContext) {
	localctx = NewLiteralExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, StrictusParserRULE_literalExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(653)
		p.Literal()
	}

	return localctx
}

// IFunctionExpressionContext is an interface to support dynamic dispatch.
type IFunctionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeAnnotationContext

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeAnnotationContext)

	// IsFunctionExpressionContext differentiates from other interfaces.
	IsFunctionExpressionContext()
}

type FunctionExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	returnType ITypeAnnotationContext
}

func NewEmptyFunctionExpressionContext() *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_functionExpression
	return p
}

func (*FunctionExpressionContext) IsFunctionExpressionContext() {}

func NewFunctionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_functionExpression

	return p
}

func (s *FunctionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionExpressionContext) GetReturnType() ITypeAnnotationContext { return s.returnType }

func (s *FunctionExpressionContext) SetReturnType(v ITypeAnnotationContext) { s.returnType = v }

func (s *FunctionExpressionContext) Fun() antlr.TerminalNode {
	return s.GetToken(StrictusParserFun, 0)
}

func (s *FunctionExpressionContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionExpressionContext) FunctionBlock() IFunctionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBlockContext)
}

func (s *FunctionExpressionContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFunctionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) FunctionExpression() (localctx IFunctionExpressionContext) {
	localctx = NewFunctionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, StrictusParserRULE_functionExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(655)
		p.Match(StrictusParserFun)
	}
	{
		p.SetState(656)
		p.ParameterList()
	}
	p.SetState(659)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserT__4 {
		{
			p.SetState(657)
			p.Match(StrictusParserT__4)
		}
		{
			p.SetState(658)

			var _x = p.TypeAnnotation()

			localctx.(*FunctionExpressionContext).returnType = _x
		}

	}
	{
		p.SetState(661)
		p.FunctionBlock()
	}

	return localctx
}

// INestedExpressionContext is an interface to support dynamic dispatch.
type INestedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNestedExpressionContext differentiates from other interfaces.
	IsNestedExpressionContext()
}

type NestedExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedExpressionContext() *NestedExpressionContext {
	var p = new(NestedExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_nestedExpression
	return p
}

func (*NestedExpressionContext) IsNestedExpressionContext() {}

func NewNestedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedExpressionContext {
	var p = new(NestedExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_nestedExpression

	return p
}

func (s *NestedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedExpressionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, 0)
}

func (s *NestedExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NestedExpressionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, 0)
}

func (s *NestedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterNestedExpression(s)
	}
}

func (s *NestedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitNestedExpression(s)
	}
}

func (s *NestedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitNestedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) NestedExpression() (localctx INestedExpressionContext) {
	localctx = NewNestedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, StrictusParserRULE_nestedExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(663)
		p.Match(StrictusParserOpenParen)
	}
	{
		p.SetState(664)
		p.Expression()
	}
	{
		p.SetState(665)
		p.Match(StrictusParserCloseParen)
	}

	return localctx
}

// IExpressionAccessContext is an interface to support dynamic dispatch.
type IExpressionAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionAccessContext differentiates from other interfaces.
	IsExpressionAccessContext()
}

type ExpressionAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAccessContext() *ExpressionAccessContext {
	var p = new(ExpressionAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_expressionAccess
	return p
}

func (*ExpressionAccessContext) IsExpressionAccessContext() {}

func NewExpressionAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAccessContext {
	var p = new(ExpressionAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_expressionAccess

	return p
}

func (s *ExpressionAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAccessContext) MemberAccess() IMemberAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberAccessContext)
}

func (s *ExpressionAccessContext) BracketExpression() IBracketExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBracketExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBracketExpressionContext)
}

func (s *ExpressionAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterExpressionAccess(s)
	}
}

func (s *ExpressionAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitExpressionAccess(s)
	}
}

func (s *ExpressionAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitExpressionAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ExpressionAccess() (localctx IExpressionAccessContext) {
	localctx = NewExpressionAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, StrictusParserRULE_expressionAccess)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(669)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(667)
			p.MemberAccess()
		}

	case StrictusParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(668)
			p.BracketExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMemberAccessContext is an interface to support dynamic dispatch.
type IMemberAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberAccessContext differentiates from other interfaces.
	IsMemberAccessContext()
}

type MemberAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberAccessContext() *MemberAccessContext {
	var p = new(MemberAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_memberAccess
	return p
}

func (*MemberAccessContext) IsMemberAccessContext() {}

func NewMemberAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberAccessContext {
	var p = new(MemberAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_memberAccess

	return p
}

func (s *MemberAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberAccessContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MemberAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterMemberAccess(s)
	}
}

func (s *MemberAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitMemberAccess(s)
	}
}

func (s *MemberAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitMemberAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) MemberAccess() (localctx IMemberAccessContext) {
	localctx = NewMemberAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, StrictusParserRULE_memberAccess)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(671)
		p.Match(StrictusParserT__11)
	}
	{
		p.SetState(672)
		p.Identifier()
	}

	return localctx
}

// IBracketExpressionContext is an interface to support dynamic dispatch.
type IBracketExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBracketExpressionContext differentiates from other interfaces.
	IsBracketExpressionContext()
}

type BracketExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketExpressionContext() *BracketExpressionContext {
	var p = new(BracketExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_bracketExpression
	return p
}

func (*BracketExpressionContext) IsBracketExpressionContext() {}

func NewBracketExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracketExpressionContext {
	var p = new(BracketExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_bracketExpression

	return p
}

func (s *BracketExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BracketExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BracketExpressionContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *BracketExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BracketExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBracketExpression(s)
	}
}

func (s *BracketExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBracketExpression(s)
	}
}

func (s *BracketExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBracketExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) BracketExpression() (localctx IBracketExpressionContext) {
	localctx = NewBracketExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, StrictusParserRULE_bracketExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(674)
		p.Match(StrictusParserT__5)
	}
	p.SetState(677)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(675)
			p.Expression()
		}

	case 2:
		{
			p.SetState(676)
			p.FullType()
		}

	}
	{
		p.SetState(679)
		p.Match(StrictusParserT__6)
	}

	return localctx
}

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_invocation
	return p
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, 0)
}

func (s *InvocationContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, 0)
}

func (s *InvocationContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *InvocationContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterInvocation(s)
	}
}

func (s *InvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitInvocation(s)
	}
}

func (s *InvocationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitInvocation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Invocation() (localctx IInvocationContext) {
	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, StrictusParserRULE_invocation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(681)
		p.Match(StrictusParserOpenParen)
	}
	p.SetState(690)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__2)|(1<<StrictusParserT__5)|(1<<StrictusParserMinus)|(1<<StrictusParserAmpersand)|(1<<StrictusParserNegate)|(1<<StrictusParserMove)|(1<<StrictusParserOpenParen))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(StrictusParserFun-38))|(1<<(StrictusParserTrue-38))|(1<<(StrictusParserFalse-38))|(1<<(StrictusParserNil-38))|(1<<(StrictusParserFrom-38))|(1<<(StrictusParserCreate-38))|(1<<(StrictusParserDestroy-38))|(1<<(StrictusParserIdentifier-38))|(1<<(StrictusParserDecimalLiteral-38))|(1<<(StrictusParserBinaryLiteral-38))|(1<<(StrictusParserOctalLiteral-38))|(1<<(StrictusParserHexadecimalLiteral-38))|(1<<(StrictusParserInvalidNumberLiteral-38))|(1<<(StrictusParserStringLiteral-38)))) != 0) {
		{
			p.SetState(682)
			p.Argument()
		}
		p.SetState(687)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(683)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(684)
				p.Argument()
			}

			p.SetState(689)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(692)
		p.Match(StrictusParserCloseParen)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 152, StrictusParserRULE_argument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(697)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(694)
			p.Identifier()
		}
		{
			p.SetState(695)
			p.Match(StrictusParserT__4)
		}

	}
	{
		p.SetState(699)
		p.Expression()
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *LiteralContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *LiteralContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *LiteralContext) DictionaryLiteral() IDictionaryLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictionaryLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDictionaryLiteralContext)
}

func (s *LiteralContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *LiteralContext) NilLiteral() INilLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INilLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INilLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 154, StrictusParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(707)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserMinus, StrictusParserDecimalLiteral, StrictusParserBinaryLiteral, StrictusParserOctalLiteral, StrictusParserHexadecimalLiteral, StrictusParserInvalidNumberLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(701)
			p.IntegerLiteral()
		}

	case StrictusParserTrue, StrictusParserFalse:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(702)
			p.BooleanLiteral()
		}

	case StrictusParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(703)
			p.ArrayLiteral()
		}

	case StrictusParserT__2:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(704)
			p.DictionaryLiteral()
		}

	case StrictusParserStringLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(705)
			p.StringLiteral()
		}

	case StrictusParserNil:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(706)
			p.NilLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) True() antlr.TerminalNode {
	return s.GetToken(StrictusParserTrue, 0)
}

func (s *BooleanLiteralContext) False() antlr.TerminalNode {
	return s.GetToken(StrictusParserFalse, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 156, StrictusParserRULE_booleanLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(709)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserTrue || _la == StrictusParserFalse) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INilLiteralContext is an interface to support dynamic dispatch.
type INilLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNilLiteralContext differentiates from other interfaces.
	IsNilLiteralContext()
}

type NilLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNilLiteralContext() *NilLiteralContext {
	var p = new(NilLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_nilLiteral
	return p
}

func (*NilLiteralContext) IsNilLiteralContext() {}

func NewNilLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NilLiteralContext {
	var p = new(NilLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_nilLiteral

	return p
}

func (s *NilLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NilLiteralContext) Nil() antlr.TerminalNode {
	return s.GetToken(StrictusParserNil, 0)
}

func (s *NilLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NilLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterNilLiteral(s)
	}
}

func (s *NilLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitNilLiteral(s)
	}
}

func (s *NilLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitNilLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) NilLiteral() (localctx INilLiteralContext) {
	localctx = NewNilLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 158, StrictusParserRULE_nilLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(711)
		p.Match(StrictusParserNil)
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserStringLiteral, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 160, StrictusParserRULE_stringLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(713)
		p.Match(StrictusParserStringLiteral)
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) PositiveIntegerLiteral() IPositiveIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPositiveIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPositiveIntegerLiteralContext)
}

func (s *IntegerLiteralContext) Minus() antlr.TerminalNode {
	return s.GetToken(StrictusParserMinus, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 162, StrictusParserRULE_integerLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(716)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserMinus {
		{
			p.SetState(715)
			p.Match(StrictusParserMinus)
		}

	}
	{
		p.SetState(718)
		p.PositiveIntegerLiteral()
	}

	return localctx
}

// IPositiveIntegerLiteralContext is an interface to support dynamic dispatch.
type IPositiveIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPositiveIntegerLiteralContext differentiates from other interfaces.
	IsPositiveIntegerLiteralContext()
}

type PositiveIntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPositiveIntegerLiteralContext() *PositiveIntegerLiteralContext {
	var p = new(PositiveIntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_positiveIntegerLiteral
	return p
}

func (*PositiveIntegerLiteralContext) IsPositiveIntegerLiteralContext() {}

func NewPositiveIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PositiveIntegerLiteralContext {
	var p = new(PositiveIntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_positiveIntegerLiteral

	return p
}

func (s *PositiveIntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *PositiveIntegerLiteralContext) CopyFrom(ctx *PositiveIntegerLiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PositiveIntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PositiveIntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryLiteralContext struct {
	*PositiveIntegerLiteralContext
}

func NewBinaryLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryLiteralContext {
	var p = new(BinaryLiteralContext)

	p.PositiveIntegerLiteralContext = NewEmptyPositiveIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PositiveIntegerLiteralContext))

	return p
}

func (s *BinaryLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLiteralContext) BinaryLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserBinaryLiteral, 0)
}

func (s *BinaryLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBinaryLiteral(s)
	}
}

func (s *BinaryLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBinaryLiteral(s)
	}
}

func (s *BinaryLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBinaryLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type OctalLiteralContext struct {
	*PositiveIntegerLiteralContext
}

func NewOctalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OctalLiteralContext {
	var p = new(OctalLiteralContext)

	p.PositiveIntegerLiteralContext = NewEmptyPositiveIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PositiveIntegerLiteralContext))

	return p
}

func (s *OctalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OctalLiteralContext) OctalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserOctalLiteral, 0)
}

func (s *OctalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterOctalLiteral(s)
	}
}

func (s *OctalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitOctalLiteral(s)
	}
}

func (s *OctalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitOctalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type InvalidNumberLiteralContext struct {
	*PositiveIntegerLiteralContext
}

func NewInvalidNumberLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvalidNumberLiteralContext {
	var p = new(InvalidNumberLiteralContext)

	p.PositiveIntegerLiteralContext = NewEmptyPositiveIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PositiveIntegerLiteralContext))

	return p
}

func (s *InvalidNumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvalidNumberLiteralContext) InvalidNumberLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserInvalidNumberLiteral, 0)
}

func (s *InvalidNumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterInvalidNumberLiteral(s)
	}
}

func (s *InvalidNumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitInvalidNumberLiteral(s)
	}
}

func (s *InvalidNumberLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitInvalidNumberLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecimalLiteralContext struct {
	*PositiveIntegerLiteralContext
}

func NewDecimalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.PositiveIntegerLiteralContext = NewEmptyPositiveIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PositiveIntegerLiteralContext))

	return p
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserDecimalLiteral, 0)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitDecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type HexadecimalLiteralContext struct {
	*PositiveIntegerLiteralContext
}

func NewHexadecimalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HexadecimalLiteralContext {
	var p = new(HexadecimalLiteralContext)

	p.PositiveIntegerLiteralContext = NewEmptyPositiveIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PositiveIntegerLiteralContext))

	return p
}

func (s *HexadecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexadecimalLiteralContext) HexadecimalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserHexadecimalLiteral, 0)
}

func (s *HexadecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterHexadecimalLiteral(s)
	}
}

func (s *HexadecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitHexadecimalLiteral(s)
	}
}

func (s *HexadecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitHexadecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PositiveIntegerLiteral() (localctx IPositiveIntegerLiteralContext) {
	localctx = NewPositiveIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 164, StrictusParserRULE_positiveIntegerLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(725)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserDecimalLiteral:
		localctx = NewDecimalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(720)
			p.Match(StrictusParserDecimalLiteral)
		}

	case StrictusParserBinaryLiteral:
		localctx = NewBinaryLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(721)
			p.Match(StrictusParserBinaryLiteral)
		}

	case StrictusParserOctalLiteral:
		localctx = NewOctalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(722)
			p.Match(StrictusParserOctalLiteral)
		}

	case StrictusParserHexadecimalLiteral:
		localctx = NewHexadecimalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(723)
			p.Match(StrictusParserHexadecimalLiteral)
		}

	case StrictusParserInvalidNumberLiteral:
		localctx = NewInvalidNumberLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(724)
			p.Match(StrictusParserInvalidNumberLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 166, StrictusParserRULE_arrayLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(727)
		p.Match(StrictusParserT__5)
	}
	p.SetState(736)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__2)|(1<<StrictusParserT__5)|(1<<StrictusParserMinus)|(1<<StrictusParserAmpersand)|(1<<StrictusParserNegate)|(1<<StrictusParserMove)|(1<<StrictusParserOpenParen))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(StrictusParserFun-38))|(1<<(StrictusParserTrue-38))|(1<<(StrictusParserFalse-38))|(1<<(StrictusParserNil-38))|(1<<(StrictusParserFrom-38))|(1<<(StrictusParserCreate-38))|(1<<(StrictusParserDestroy-38))|(1<<(StrictusParserIdentifier-38))|(1<<(StrictusParserDecimalLiteral-38))|(1<<(StrictusParserBinaryLiteral-38))|(1<<(StrictusParserOctalLiteral-38))|(1<<(StrictusParserHexadecimalLiteral-38))|(1<<(StrictusParserInvalidNumberLiteral-38))|(1<<(StrictusParserStringLiteral-38)))) != 0) {
		{
			p.SetState(728)
			p.Expression()
		}
		p.SetState(733)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(729)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(730)
				p.Expression()
			}

			p.SetState(735)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(738)
		p.Match(StrictusParserT__6)
	}

	return localctx
}

// IDictionaryLiteralContext is an interface to support dynamic dispatch.
type IDictionaryLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDictionaryLiteralContext differentiates from other interfaces.
	IsDictionaryLiteralContext()
}

type DictionaryLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictionaryLiteralContext() *DictionaryLiteralContext {
	var p = new(DictionaryLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_dictionaryLiteral
	return p
}

func (*DictionaryLiteralContext) IsDictionaryLiteralContext() {}

func NewDictionaryLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictionaryLiteralContext {
	var p = new(DictionaryLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_dictionaryLiteral

	return p
}

func (s *DictionaryLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DictionaryLiteralContext) AllDictionaryEntry() []IDictionaryEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDictionaryEntryContext)(nil)).Elem())
	var tst = make([]IDictionaryEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDictionaryEntryContext)
		}
	}

	return tst
}

func (s *DictionaryLiteralContext) DictionaryEntry(i int) IDictionaryEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictionaryEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDictionaryEntryContext)
}

func (s *DictionaryLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictionaryLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictionaryLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterDictionaryLiteral(s)
	}
}

func (s *DictionaryLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitDictionaryLiteral(s)
	}
}

func (s *DictionaryLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitDictionaryLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) DictionaryLiteral() (localctx IDictionaryLiteralContext) {
	localctx = NewDictionaryLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 168, StrictusParserRULE_dictionaryLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(740)
		p.Match(StrictusParserT__2)
	}
	p.SetState(749)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__2)|(1<<StrictusParserT__5)|(1<<StrictusParserMinus)|(1<<StrictusParserAmpersand)|(1<<StrictusParserNegate)|(1<<StrictusParserMove)|(1<<StrictusParserOpenParen))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(StrictusParserFun-38))|(1<<(StrictusParserTrue-38))|(1<<(StrictusParserFalse-38))|(1<<(StrictusParserNil-38))|(1<<(StrictusParserFrom-38))|(1<<(StrictusParserCreate-38))|(1<<(StrictusParserDestroy-38))|(1<<(StrictusParserIdentifier-38))|(1<<(StrictusParserDecimalLiteral-38))|(1<<(StrictusParserBinaryLiteral-38))|(1<<(StrictusParserOctalLiteral-38))|(1<<(StrictusParserHexadecimalLiteral-38))|(1<<(StrictusParserInvalidNumberLiteral-38))|(1<<(StrictusParserStringLiteral-38)))) != 0) {
		{
			p.SetState(741)
			p.DictionaryEntry()
		}
		p.SetState(746)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(742)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(743)
				p.DictionaryEntry()
			}

			p.SetState(748)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(751)
		p.Match(StrictusParserT__3)
	}

	return localctx
}

// IDictionaryEntryContext is an interface to support dynamic dispatch.
type IDictionaryEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key rule contexts.
	GetKey() IExpressionContext

	// GetValue returns the value rule contexts.
	GetValue() IExpressionContext

	// SetKey sets the key rule contexts.
	SetKey(IExpressionContext)

	// SetValue sets the value rule contexts.
	SetValue(IExpressionContext)

	// IsDictionaryEntryContext differentiates from other interfaces.
	IsDictionaryEntryContext()
}

type DictionaryEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    IExpressionContext
	value  IExpressionContext
}

func NewEmptyDictionaryEntryContext() *DictionaryEntryContext {
	var p = new(DictionaryEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_dictionaryEntry
	return p
}

func (*DictionaryEntryContext) IsDictionaryEntryContext() {}

func NewDictionaryEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictionaryEntryContext {
	var p = new(DictionaryEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_dictionaryEntry

	return p
}

func (s *DictionaryEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *DictionaryEntryContext) GetKey() IExpressionContext { return s.key }

func (s *DictionaryEntryContext) GetValue() IExpressionContext { return s.value }

func (s *DictionaryEntryContext) SetKey(v IExpressionContext) { s.key = v }

func (s *DictionaryEntryContext) SetValue(v IExpressionContext) { s.value = v }

func (s *DictionaryEntryContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DictionaryEntryContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DictionaryEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictionaryEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictionaryEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterDictionaryEntry(s)
	}
}

func (s *DictionaryEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitDictionaryEntry(s)
	}
}

func (s *DictionaryEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitDictionaryEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) DictionaryEntry() (localctx IDictionaryEntryContext) {
	localctx = NewDictionaryEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 170, StrictusParserRULE_dictionaryEntry)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(753)

		var _x = p.Expression()

		localctx.(*DictionaryEntryContext).key = _x
	}
	{
		p.SetState(754)
		p.Match(StrictusParserT__4)
	}
	{
		p.SetState(755)

		var _x = p.Expression()

		localctx.(*DictionaryEntryContext).value = _x
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *IdentifierContext) From() antlr.TerminalNode {
	return s.GetToken(StrictusParserFrom, 0)
}

func (s *IdentifierContext) Create() antlr.TerminalNode {
	return s.GetToken(StrictusParserCreate, 0)
}

func (s *IdentifierContext) Destroy() antlr.TerminalNode {
	return s.GetToken(StrictusParserDestroy, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 172, StrictusParserRULE_identifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(757)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(StrictusParserFrom-57))|(1<<(StrictusParserCreate-57))|(1<<(StrictusParserDestroy-57))|(1<<(StrictusParserIdentifier-57)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(StrictusParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 174, StrictusParserRULE_eos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(763)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(759)
			p.Match(StrictusParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(760)
			p.Match(StrictusParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(761)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "p.lineTerminatorAhead()", ""))
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(762)

		if !(p.GetTokenStream().LT(1).GetText() == "}") {
			panic(antlr.NewFailedPredicateException(p, "p.GetTokenStream().LT(1).GetText() == \"}\"", ""))
		}

	}

	return localctx
}

func (p *StrictusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
		var t *SpecialFunctionDeclarationContext = nil
		if localctx != nil {
			t = localctx.(*SpecialFunctionDeclarationContext)
		}
		return p.SpecialFunctionDeclaration_Sempred(t, predIndex)

	case 14:
		var t *FunctionDeclarationContext = nil
		if localctx != nil {
			t = localctx.(*FunctionDeclarationContext)
		}
		return p.FunctionDeclaration_Sempred(t, predIndex)

	case 19:
		var t *FullTypeContext = nil
		if localctx != nil {
			t = localctx.(*FullTypeContext)
		}
		return p.FullType_Sempred(t, predIndex)

	case 34:
		var t *ReturnStatementContext = nil
		if localctx != nil {
			t = localctx.(*ReturnStatementContext)
		}
		return p.ReturnStatement_Sempred(t, predIndex)

	case 46:
		var t *OrExpressionContext = nil
		if localctx != nil {
			t = localctx.(*OrExpressionContext)
		}
		return p.OrExpression_Sempred(t, predIndex)

	case 47:
		var t *AndExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AndExpressionContext)
		}
		return p.AndExpression_Sempred(t, predIndex)

	case 48:
		var t *EqualityExpressionContext = nil
		if localctx != nil {
			t = localctx.(*EqualityExpressionContext)
		}
		return p.EqualityExpression_Sempred(t, predIndex)

	case 49:
		var t *RelationalExpressionContext = nil
		if localctx != nil {
			t = localctx.(*RelationalExpressionContext)
		}
		return p.RelationalExpression_Sempred(t, predIndex)

	case 51:
		var t *FailableDowncastingExpressionContext = nil
		if localctx != nil {
			t = localctx.(*FailableDowncastingExpressionContext)
		}
		return p.FailableDowncastingExpression_Sempred(t, predIndex)

	case 52:
		var t *ConcatenatingExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ConcatenatingExpressionContext)
		}
		return p.ConcatenatingExpression_Sempred(t, predIndex)

	case 53:
		var t *AdditiveExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AdditiveExpressionContext)
		}
		return p.AdditiveExpression_Sempred(t, predIndex)

	case 54:
		var t *MultiplicativeExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MultiplicativeExpressionContext)
		}
		return p.MultiplicativeExpression_Sempred(t, predIndex)

	case 87:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *StrictusParser) SpecialFunctionDeclaration_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return !localctx.(*SpecialFunctionDeclarationContext).functionBlockRequired || localctx.(*SpecialFunctionDeclarationContext).b != nil

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) FunctionDeclaration_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return !localctx.(*FunctionDeclarationContext).functionBlockRequired || localctx.(*FunctionDeclarationContext).b != nil

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) FullType_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.noWhitespace()

	case 3:
		return p.noWhitespace()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) ReturnStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return !p.lineTerminatorAhead()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) OrExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) AndExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) EqualityExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) RelationalExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) FailableDowncastingExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) ConcatenatingExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) AdditiveExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) MultiplicativeExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 13:
		return p.lineTerminatorAhead()

	case 14:
		return p.GetTokenStream().LT(1).GetText() == "}"

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

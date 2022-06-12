typedef struct {
	uint8_t firstOffset[3];
	uint8_t unused;
	uint8_t lastOffset[3];
	uint8_t unused;
} oki_entry;

typedef struct {
	uint8_t unused[8]
	oki_entry entries[127]
	uint8_t payload[0x3FC00]
}
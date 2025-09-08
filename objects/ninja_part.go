package objects

func (this *Ninja) PartHead() int16 {
	itemHeadHide := this.ItemBodyHide[0]
	if itemHeadHide != nil {
		switch itemHeadHide.ID {
		case 711:
			return 223
		case 774:
			return 267
		case 786:
			if this.Gender == 0 {
				return 273
			} else {
				return 270
			}
		case 787:
			if this.Gender == 0 {
				return 279
			} else {
				return 276
			}
		}
	}
	if this.ItemBody[11] == nil {
		return this.Head
	}
	if this.ItemBody[2] != nil && (this.ItemBody[2].ID == 795 || this.ItemBody[2].ID == 796) {
		return -1
	}
	ItemBody := this.ItemBody
	if ItemBody[11].ID >= 813 && ItemBody[11].ID <= 818 {
		return -1
	}
	if this.ItemBody[11].ID == 541 {
		return 185
	}
	if this.ItemBody[11].ID == 542 {
		return 188
	}
	if this.ItemBody[11].ID == 745 {
		return 264
	}
	if this.ItemBody[11].ID == 774 {
		return 267
	}
	if this.ItemBody[11].ID == 786 {
		return 270
	}
	if this.ItemBody[11].ID == 787 {
		return 276
	}
	if this.ItemBody[11].ID == 853 {
		return 273
	}
	if this.ItemBody[11].ID == 854 {
		return 279
	}
	if this.ItemBody[11].ID == 711 {
		return 223
	}
	return this.ItemBody[11].GetData().Part
}

func (this *Ninja) Weapon() int16 {
	if this.ItemBody[1] != nil {
		return this.ItemBody[1].GetData().Part
	}
	return -1
}

func (this *Ninja) PartBody() int16 {
	head := this.PartHead()
	switch head {
	case 223:
		return 224
	case 226:
		return 227
	case 280:
		return 281
	case 185:
		return 186
	case 188:
		return 189
	case 258:
		return 259
	case 264:
		return 265
	case 267:
		return 268
	case 270:
		return 271
	case 273:
		return 274
	case 279:
		return 280
	}
	ao := this.ItemBody[2]
	if ao != nil {
		if ao.ID == 795 || ao.ID == 796 {
			return -1
		}
		return ao.GetData().Part
	}
	return -1
}

func (this *Ninja) PartLeg() int16 {
	head := this.PartHead()
	switch head {
	case 223:
		return 225
	case 226:
		return 228
	case 270:
		return 272
	case 185:
		return 187
	case 188:
		return 190
	case 258:
		return 260
	case 264:
		return 266
	case 267:
		return 269
	case 276:
		return 278
	case 273:
		return 275
	case 279:
		return 281
	}
	quan := this.ItemBody[6]
	if quan != nil {
		if this.ItemBody[2] != nil && (this.ItemBody[2].ID == 795 || this.ItemBody[2].ID == 796) {
			return -1
		}
		return quan.GetData().Part
	}
	return -1
}

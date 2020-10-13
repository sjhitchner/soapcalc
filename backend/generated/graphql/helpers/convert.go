// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package helpers

import (
	"github.com/sjhitchner/soapcalc/backend/generated/db/models"
	gmodels "github.com/sjhitchner/soapcalc/backend/generated/graphql/models"
	null "github.com/volatiletech/null/v8"
	"github.com/web-ridge/utils-go/boilergql"
)

func AdditiveWithUintID(id uint) *gmodels.Additive {
	return &gmodels.Additive{
		ID: AdditiveIDToGraphQL(id),
	}
}

func AdditiveWithIntID(id int) *gmodels.Additive {
	return AdditiveWithUintID(uint(id))
}

func AdditiveWithNullDotUintID(id null.Uint) *gmodels.Additive {
	return AdditiveWithUintID(id.Uint)
}

func AdditiveWithNullDotIntID(id null.Int) *gmodels.Additive {
	return AdditiveWithUintID(uint(id.Int))
}

func AdditivesToGraphQL(am []*models.Additive) []*gmodels.Additive {
	ar := make([]*gmodels.Additive, len(am))
	for i, m := range am {
		ar[i] = AdditiveToGraphQL(m)
	}
	return ar
}

func AdditiveIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.Additive)
}

func AdditiveToGraphQL(m *models.Additive) *gmodels.Additive {
	if m == nil {
		return nil
	}

	r := &gmodels.Additive{
		ID:        AdditiveIDToGraphQL(uint(m.ID)),
		Name:      m.Name,
		Note:      m.Note,
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	return r
}

func AdditiveID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func AdditiveIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func AdditiveInventoryWithUintID(id uint) *gmodels.AdditiveInventory {
	return &gmodels.AdditiveInventory{
		ID: AdditiveInventoryIDToGraphQL(id),
	}
}

func AdditiveInventoryWithIntID(id int) *gmodels.AdditiveInventory {
	return AdditiveInventoryWithUintID(uint(id))
}

func AdditiveInventoryWithNullDotUintID(id null.Uint) *gmodels.AdditiveInventory {
	return AdditiveInventoryWithUintID(id.Uint)
}

func AdditiveInventoryWithNullDotIntID(id null.Int) *gmodels.AdditiveInventory {
	return AdditiveInventoryWithUintID(uint(id.Int))
}

func AdditiveInventoriesToGraphQL(am []*models.AdditiveInventory) []*gmodels.AdditiveInventory {
	ar := make([]*gmodels.AdditiveInventory, len(am))
	for i, m := range am {
		ar[i] = AdditiveInventoryToGraphQL(m)
	}
	return ar
}

func AdditiveInventoryIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.AdditiveInventory)
}

func AdditiveInventoryToGraphQL(m *models.AdditiveInventory) *gmodels.AdditiveInventory {
	if m == nil {
		return nil
	}

	r := &gmodels.AdditiveInventory{
		ID:           AdditiveInventoryIDToGraphQL(uint(m.ID)),
		PurchaseDate: boilergql.TimeDotTimeToInt(m.PurchaseDate),
		ExpiryDate:   boilergql.TimeDotTimeToInt(m.ExpiryDate),
		Cost:         m.Cost,
		Weight:       m.Weight,
		UpdatedAt:    boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:    boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		CreatedAt:    boilergql.TimeDotTimeToInt(m.CreatedAt),
	}

	if boilergql.IntIsFilled(m.AdditiveID) {
		if m.R != nil && m.R.Additive != nil {
			r.Additive = AdditiveToGraphQL(m.R.Additive)
		} else {
			r.Additive = AdditiveWithIntID(m.AdditiveID)
		}
	}

	if boilergql.IntIsFilled(m.SupplierID) {
		if m.R != nil && m.R.Supplier != nil {
			r.Supplier = SupplierToGraphQL(m.R.Supplier)
		} else {
			r.Supplier = SupplierWithIntID(m.SupplierID)
		}
	}

	return r
}

func AdditiveInventoryID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func AdditiveInventoryIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func FragranceWithUintID(id uint) *gmodels.Fragrance {
	return &gmodels.Fragrance{
		ID: FragranceIDToGraphQL(id),
	}
}

func FragranceWithIntID(id int) *gmodels.Fragrance {
	return FragranceWithUintID(uint(id))
}

func FragranceWithNullDotUintID(id null.Uint) *gmodels.Fragrance {
	return FragranceWithUintID(id.Uint)
}

func FragranceWithNullDotIntID(id null.Int) *gmodels.Fragrance {
	return FragranceWithUintID(uint(id.Int))
}

func FragrancesToGraphQL(am []*models.Fragrance) []*gmodels.Fragrance {
	ar := make([]*gmodels.Fragrance, len(am))
	for i, m := range am {
		ar[i] = FragranceToGraphQL(m)
	}
	return ar
}

func FragranceIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.Fragrance)
}

func FragranceToGraphQL(m *models.Fragrance) *gmodels.Fragrance {
	if m == nil {
		return nil
	}

	r := &gmodels.Fragrance{
		ID:        FragranceIDToGraphQL(uint(m.ID)),
		Name:      m.Name,
		Note:      m.Note,
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
	}

	return r
}

func FragranceID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func FragranceIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func FragranceInventoryWithUintID(id uint) *gmodels.FragranceInventory {
	return &gmodels.FragranceInventory{
		ID: FragranceInventoryIDToGraphQL(id),
	}
}

func FragranceInventoryWithIntID(id int) *gmodels.FragranceInventory {
	return FragranceInventoryWithUintID(uint(id))
}

func FragranceInventoryWithNullDotUintID(id null.Uint) *gmodels.FragranceInventory {
	return FragranceInventoryWithUintID(id.Uint)
}

func FragranceInventoryWithNullDotIntID(id null.Int) *gmodels.FragranceInventory {
	return FragranceInventoryWithUintID(uint(id.Int))
}

func FragranceInventoriesToGraphQL(am []*models.FragranceInventory) []*gmodels.FragranceInventory {
	ar := make([]*gmodels.FragranceInventory, len(am))
	for i, m := range am {
		ar[i] = FragranceInventoryToGraphQL(m)
	}
	return ar
}

func FragranceInventoryIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.FragranceInventory)
}

func FragranceInventoryToGraphQL(m *models.FragranceInventory) *gmodels.FragranceInventory {
	if m == nil {
		return nil
	}

	r := &gmodels.FragranceInventory{
		ID:           FragranceInventoryIDToGraphQL(uint(m.ID)),
		PurchaseDate: boilergql.TimeDotTimeToInt(m.PurchaseDate),
		ExpiryDate:   boilergql.TimeDotTimeToInt(m.ExpiryDate),
		Cost:         m.Cost,
		Weight:       m.Weight,
		DeletedAt:    boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		CreatedAt:    boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:    boilergql.TimeDotTimeToInt(m.UpdatedAt),
	}

	if boilergql.IntIsFilled(m.FragranceID) {
		if m.R != nil && m.R.Fragrance != nil {
			r.Fragrance = FragranceToGraphQL(m.R.Fragrance)
		} else {
			r.Fragrance = FragranceWithIntID(m.FragranceID)
		}
	}

	if boilergql.IntIsFilled(m.SupplierID) {
		if m.R != nil && m.R.Supplier != nil {
			r.Supplier = SupplierToGraphQL(m.R.Supplier)
		} else {
			r.Supplier = SupplierWithIntID(m.SupplierID)
		}
	}

	return r
}

func FragranceInventoryID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func FragranceInventoryIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func LipidWithUintID(id uint) *gmodels.Lipid {
	return &gmodels.Lipid{
		ID: LipidIDToGraphQL(id),
	}
}

func LipidWithIntID(id int) *gmodels.Lipid {
	return LipidWithUintID(uint(id))
}

func LipidWithNullDotUintID(id null.Uint) *gmodels.Lipid {
	return LipidWithUintID(id.Uint)
}

func LipidWithNullDotIntID(id null.Int) *gmodels.Lipid {
	return LipidWithUintID(uint(id.Int))
}

func LipidsToGraphQL(am []*models.Lipid) []*gmodels.Lipid {
	ar := make([]*gmodels.Lipid, len(am))
	for i, m := range am {
		ar[i] = LipidToGraphQL(m)
	}
	return ar
}

func LipidIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.Lipid)
}

func LipidToGraphQL(m *models.Lipid) *gmodels.Lipid {
	if m == nil {
		return nil
	}

	r := &gmodels.Lipid{
		ID:           LipidIDToGraphQL(uint(m.ID)),
		Name:         m.Name,
		Lauric:       m.Lauric,
		Myristic:     m.Myristic,
		Palmitic:     m.Palmitic,
		Stearic:      m.Stearic,
		Ricinoleic:   m.Ricinoleic,
		Oleic:        m.Oleic,
		Linoleic:     m.Linoleic,
		Linolenic:    m.Linolenic,
		Hardness:     m.Hardness,
		Cleansing:    m.Cleansing,
		Conditioning: m.Conditioning,
		Bubbly:       m.Bubbly,
		Creamy:       m.Creamy,
		Iodine:       m.Iodine,
		Ins:          m.Ins,
		InciName:     m.InciName,
		Family:       m.Family,
		Naoh:         m.Naoh,
		DeletedAt:    boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		CreatedAt:    boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:    boilergql.TimeDotTimeToInt(m.UpdatedAt),
	}

	return r
}

func LipidID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func LipidIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func LipidInventoryWithUintID(id uint) *gmodels.LipidInventory {
	return &gmodels.LipidInventory{
		ID: LipidInventoryIDToGraphQL(id),
	}
}

func LipidInventoryWithIntID(id int) *gmodels.LipidInventory {
	return LipidInventoryWithUintID(uint(id))
}

func LipidInventoryWithNullDotUintID(id null.Uint) *gmodels.LipidInventory {
	return LipidInventoryWithUintID(id.Uint)
}

func LipidInventoryWithNullDotIntID(id null.Int) *gmodels.LipidInventory {
	return LipidInventoryWithUintID(uint(id.Int))
}

func LipidInventoriesToGraphQL(am []*models.LipidInventory) []*gmodels.LipidInventory {
	ar := make([]*gmodels.LipidInventory, len(am))
	for i, m := range am {
		ar[i] = LipidInventoryToGraphQL(m)
	}
	return ar
}

func LipidInventoryIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.LipidInventory)
}

func LipidInventoryToGraphQL(m *models.LipidInventory) *gmodels.LipidInventory {
	if m == nil {
		return nil
	}

	r := &gmodels.LipidInventory{
		ID:            LipidInventoryIDToGraphQL(uint(m.ID)),
		PurchaseDate:  boilergql.TimeDotTimeToInt(m.PurchaseDate),
		ExpiryDate:    boilergql.TimeDotTimeToInt(m.ExpiryDate),
		Cost:          m.Cost,
		Weight:        m.Weight,
		Sap:           m.Sap,
		Naoh:          m.Naoh,
		Koh:           m.Koh,
		GramsPerLiter: m.GramsPerLiter, Lipid: LipidIDToGraphQL(uint(m.LipidID)),
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
	}

	if boilergql.IntIsFilled(m.LipidID) {
		if m.R != nil && m.R.Lipid != nil {
			r.Lipid = LipidToGraphQL(m.R.Lipid)
		} else {
			r.Lipid = LipidWithIntID(m.LipidID)
		}
	}

	if boilergql.IntIsFilled(m.SupplierID) {
		if m.R != nil && m.R.Supplier != nil {
			r.Supplier = SupplierToGraphQL(m.R.Supplier)
		} else {
			r.Supplier = SupplierWithIntID(m.SupplierID)
		}
	}

	return r
}

func LipidInventoryID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func LipidInventoryIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func LyeWithUintID(id uint) *gmodels.Lye {
	return &gmodels.Lye{
		ID: LyeIDToGraphQL(id),
	}
}

func LyeWithIntID(id int) *gmodels.Lye {
	return LyeWithUintID(uint(id))
}

func LyeWithNullDotUintID(id null.Uint) *gmodels.Lye {
	return LyeWithUintID(id.Uint)
}

func LyeWithNullDotIntID(id null.Int) *gmodels.Lye {
	return LyeWithUintID(uint(id.Int))
}

func LyesToGraphQL(am []*models.Lye) []*gmodels.Lye {
	ar := make([]*gmodels.Lye, len(am))
	for i, m := range am {
		ar[i] = LyeToGraphQL(m)
	}
	return ar
}

func LyeIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.Lye)
}

func LyeToGraphQL(m *models.Lye) *gmodels.Lye {
	if m == nil {
		return nil
	}

	r := &gmodels.Lye{
		ID:        LyeIDToGraphQL(uint(m.ID)),
		Kind:      m.Kind,
		Name:      m.Name,
		Note:      m.Note,
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
	}

	return r
}

func LyeID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func LyeIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func LyeInventoryWithUintID(id uint) *gmodels.LyeInventory {
	return &gmodels.LyeInventory{
		ID: LyeInventoryIDToGraphQL(id),
	}
}

func LyeInventoryWithIntID(id int) *gmodels.LyeInventory {
	return LyeInventoryWithUintID(uint(id))
}

func LyeInventoryWithNullDotUintID(id null.Uint) *gmodels.LyeInventory {
	return LyeInventoryWithUintID(id.Uint)
}

func LyeInventoryWithNullDotIntID(id null.Int) *gmodels.LyeInventory {
	return LyeInventoryWithUintID(uint(id.Int))
}

func LyeInventoriesToGraphQL(am []*models.LyeInventory) []*gmodels.LyeInventory {
	ar := make([]*gmodels.LyeInventory, len(am))
	for i, m := range am {
		ar[i] = LyeInventoryToGraphQL(m)
	}
	return ar
}

func LyeInventoryIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.LyeInventory)
}

func LyeInventoryToGraphQL(m *models.LyeInventory) *gmodels.LyeInventory {
	if m == nil {
		return nil
	}

	r := &gmodels.LyeInventory{
		ID:            LyeInventoryIDToGraphQL(uint(m.ID)),
		PurchaseDate:  boilergql.TimeDotTimeToInt(m.PurchaseDate),
		ExpiryDate:    boilergql.TimeDotTimeToInt(m.ExpiryDate),
		Cost:          m.Cost,
		Weight:        m.Weight,
		Concentration: m.Concentration,
		DeletedAt:     boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		CreatedAt:     boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:     boilergql.TimeDotTimeToInt(m.UpdatedAt),
	}

	if boilergql.IntIsFilled(m.LyeID) {
		if m.R != nil && m.R.Lye != nil {
			r.Lye = LyeToGraphQL(m.R.Lye)
		} else {
			r.Lye = LyeWithIntID(m.LyeID)
		}
	}

	if boilergql.IntIsFilled(m.SupplierID) {
		if m.R != nil && m.R.Supplier != nil {
			r.Supplier = SupplierToGraphQL(m.R.Supplier)
		} else {
			r.Supplier = SupplierWithIntID(m.SupplierID)
		}
	}

	return r
}

func LyeInventoryID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func LyeInventoryIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeWithUintID(id uint) *gmodels.Recipe {
	return &gmodels.Recipe{
		ID: RecipeIDToGraphQL(id),
	}
}

func RecipeWithIntID(id int) *gmodels.Recipe {
	return RecipeWithUintID(uint(id))
}

func RecipeWithNullDotUintID(id null.Uint) *gmodels.Recipe {
	return RecipeWithUintID(id.Uint)
}

func RecipeWithNullDotIntID(id null.Int) *gmodels.Recipe {
	return RecipeWithUintID(uint(id.Int))
}

func RecipesToGraphQL(am []*models.Recipe) []*gmodels.Recipe {
	ar := make([]*gmodels.Recipe, len(am))
	for i, m := range am {
		ar[i] = RecipeToGraphQL(m)
	}
	return ar
}

func RecipeIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.Recipe)
}

func RecipeToGraphQL(m *models.Recipe) *gmodels.Recipe {
	if m == nil {
		return nil
	}

	r := &gmodels.Recipe{
		ID:        RecipeIDToGraphQL(uint(m.ID)),
		Name:      m.Name,
		Note:      m.Note,
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
	}

	return r
}

func RecipeID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeAdditiveWithUintID(id uint) *gmodels.RecipeAdditive {
	return &gmodels.RecipeAdditive{
		ID: RecipeAdditiveIDToGraphQL(id),
	}
}

func RecipeAdditiveWithIntID(id int) *gmodels.RecipeAdditive {
	return RecipeAdditiveWithUintID(uint(id))
}

func RecipeAdditiveWithNullDotUintID(id null.Uint) *gmodels.RecipeAdditive {
	return RecipeAdditiveWithUintID(id.Uint)
}

func RecipeAdditiveWithNullDotIntID(id null.Int) *gmodels.RecipeAdditive {
	return RecipeAdditiveWithUintID(uint(id.Int))
}

func RecipeAdditivesToGraphQL(am []*models.RecipeAdditive) []*gmodels.RecipeAdditive {
	ar := make([]*gmodels.RecipeAdditive, len(am))
	for i, m := range am {
		ar[i] = RecipeAdditiveToGraphQL(m)
	}
	return ar
}

func RecipeAdditiveIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeAdditive)
}

func RecipeAdditiveToGraphQL(m *models.RecipeAdditive) *gmodels.RecipeAdditive {
	if m == nil {
		return nil
	}

	r := &gmodels.RecipeAdditive{
		ID:         RecipeAdditiveIDToGraphQL(uint(m.ID)),
		Percentage: m.Percentage,
		DeletedAt:  boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		CreatedAt:  boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:  boilergql.TimeDotTimeToInt(m.UpdatedAt),
	}

	if boilergql.IntIsFilled(m.AdditiveID) {
		if m.R != nil && m.R.Additive != nil {
			r.Additive = AdditiveToGraphQL(m.R.Additive)
		} else {
			r.Additive = AdditiveWithIntID(m.AdditiveID)
		}
	}

	if boilergql.IntIsFilled(m.RecipeID) {
		if m.R != nil && m.R.Recipe != nil {
			r.Recipe = RecipeToGraphQL(m.R.Recipe)
		} else {
			r.Recipe = RecipeWithIntID(m.RecipeID)
		}
	}

	return r
}

func RecipeAdditiveID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeAdditiveIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeBatchWithUintID(id uint) *gmodels.RecipeBatch {
	return &gmodels.RecipeBatch{
		ID: RecipeBatchIDToGraphQL(id),
	}
}

func RecipeBatchWithIntID(id int) *gmodels.RecipeBatch {
	return RecipeBatchWithUintID(uint(id))
}

func RecipeBatchWithNullDotUintID(id null.Uint) *gmodels.RecipeBatch {
	return RecipeBatchWithUintID(id.Uint)
}

func RecipeBatchWithNullDotIntID(id null.Int) *gmodels.RecipeBatch {
	return RecipeBatchWithUintID(uint(id.Int))
}

func RecipeBatchesToGraphQL(am []*models.RecipeBatch) []*gmodels.RecipeBatch {
	ar := make([]*gmodels.RecipeBatch, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchToGraphQL(m)
	}
	return ar
}

func RecipeBatchIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeBatch)
}

func RecipeBatchToGraphQL(m *models.RecipeBatch) *gmodels.RecipeBatch {
	if m == nil {
		return nil
	}

	r := &gmodels.RecipeBatch{
		ID:               RecipeBatchIDToGraphQL(uint(m.ID)),
		Tag:              m.Tag,
		ProductionDate:   boilergql.TimeDotTimeToInt(m.ProductionDate),
		SellableDate:     boilergql.TimeDotTimeToInt(m.SellableDate),
		Note:             m.Note,
		LipidWeight:      m.LipidWeight,
		ProductionWeight: m.ProductionWeight,
		CuredWeight:      m.CuredWeight,
		DeletedAt:        boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		UpdatedAt:        boilergql.TimeDotTimeToInt(m.UpdatedAt),
		CreatedAt:        boilergql.TimeDotTimeToInt(m.CreatedAt),
	}

	if boilergql.IntIsFilled(m.RecipeID) {
		if m.R != nil && m.R.Recipe != nil {
			r.Recipe = RecipeToGraphQL(m.R.Recipe)
		} else {
			r.Recipe = RecipeWithIntID(m.RecipeID)
		}
	}

	return r
}

func RecipeBatchID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeBatchIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeBatchAdditiveWithUintID(id uint) *gmodels.RecipeBatchAdditive {
	return &gmodels.RecipeBatchAdditive{
		ID: RecipeBatchAdditiveIDToGraphQL(id),
	}
}

func RecipeBatchAdditiveWithIntID(id int) *gmodels.RecipeBatchAdditive {
	return RecipeBatchAdditiveWithUintID(uint(id))
}

func RecipeBatchAdditiveWithNullDotUintID(id null.Uint) *gmodels.RecipeBatchAdditive {
	return RecipeBatchAdditiveWithUintID(id.Uint)
}

func RecipeBatchAdditiveWithNullDotIntID(id null.Int) *gmodels.RecipeBatchAdditive {
	return RecipeBatchAdditiveWithUintID(uint(id.Int))
}

func RecipeBatchAdditivesToGraphQL(am []*models.RecipeBatchAdditive) []*gmodels.RecipeBatchAdditive {
	ar := make([]*gmodels.RecipeBatchAdditive, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchAdditiveToGraphQL(m)
	}
	return ar
}

func RecipeBatchAdditiveIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeBatchAdditive)
}

func RecipeBatchAdditiveToGraphQL(m *models.RecipeBatchAdditive) *gmodels.RecipeBatchAdditive {
	if m == nil {
		return nil
	}

	r := &gmodels.RecipeBatchAdditive{
		ID:        RecipeBatchAdditiveIDToGraphQL(uint(m.ID)),
		Weight:    m.Weight,
		Cost:      m.Cost,
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.AdditiveID) {
		if m.R != nil && m.R.Additive != nil {
			r.Additive = AdditiveToGraphQL(m.R.Additive)
		} else {
			r.Additive = AdditiveWithIntID(m.AdditiveID)
		}
	}

	if boilergql.IntIsFilled(m.BatchID) {
		if m.R != nil && m.R.Batch != nil {
			r.Batch = RecipeBatchToGraphQL(m.R.Batch)
		} else {
			r.Batch = RecipeBatchWithIntID(m.BatchID)
		}
	}

	return r
}

func RecipeBatchAdditiveID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeBatchAdditiveIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeBatchFragranceWithUintID(id uint) *gmodels.RecipeBatchFragrance {
	return &gmodels.RecipeBatchFragrance{
		ID: RecipeBatchFragranceIDToGraphQL(id),
	}
}

func RecipeBatchFragranceWithIntID(id int) *gmodels.RecipeBatchFragrance {
	return RecipeBatchFragranceWithUintID(uint(id))
}

func RecipeBatchFragranceWithNullDotUintID(id null.Uint) *gmodels.RecipeBatchFragrance {
	return RecipeBatchFragranceWithUintID(id.Uint)
}

func RecipeBatchFragranceWithNullDotIntID(id null.Int) *gmodels.RecipeBatchFragrance {
	return RecipeBatchFragranceWithUintID(uint(id.Int))
}

func RecipeBatchFragrancesToGraphQL(am []*models.RecipeBatchFragrance) []*gmodels.RecipeBatchFragrance {
	ar := make([]*gmodels.RecipeBatchFragrance, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchFragranceToGraphQL(m)
	}
	return ar
}

func RecipeBatchFragranceIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeBatchFragrance)
}

func RecipeBatchFragranceToGraphQL(m *models.RecipeBatchFragrance) *gmodels.RecipeBatchFragrance {
	if m == nil {
		return nil
	}

	r := &gmodels.RecipeBatchFragrance{
		ID:        RecipeBatchFragranceIDToGraphQL(uint(m.ID)),
		Weight:    m.Weight,
		Cost:      m.Cost,
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
	}

	if boilergql.IntIsFilled(m.FragranceID) {
		if m.R != nil && m.R.Fragrance != nil {
			r.Fragrance = FragranceToGraphQL(m.R.Fragrance)
		} else {
			r.Fragrance = FragranceWithIntID(m.FragranceID)
		}
	}

	if boilergql.IntIsFilled(m.BatchID) {
		if m.R != nil && m.R.Batch != nil {
			r.Batch = RecipeBatchToGraphQL(m.R.Batch)
		} else {
			r.Batch = RecipeBatchWithIntID(m.BatchID)
		}
	}

	return r
}

func RecipeBatchFragranceID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeBatchFragranceIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeBatchLipidWithUintID(id uint) *gmodels.RecipeBatchLipid {
	return &gmodels.RecipeBatchLipid{
		ID: RecipeBatchLipidIDToGraphQL(id),
	}
}

func RecipeBatchLipidWithIntID(id int) *gmodels.RecipeBatchLipid {
	return RecipeBatchLipidWithUintID(uint(id))
}

func RecipeBatchLipidWithNullDotUintID(id null.Uint) *gmodels.RecipeBatchLipid {
	return RecipeBatchLipidWithUintID(id.Uint)
}

func RecipeBatchLipidWithNullDotIntID(id null.Int) *gmodels.RecipeBatchLipid {
	return RecipeBatchLipidWithUintID(uint(id.Int))
}

func RecipeBatchLipidsToGraphQL(am []*models.RecipeBatchLipid) []*gmodels.RecipeBatchLipid {
	ar := make([]*gmodels.RecipeBatchLipid, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchLipidToGraphQL(m)
	}
	return ar
}

func RecipeBatchLipidIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeBatchLipid)
}

func RecipeBatchLipidToGraphQL(m *models.RecipeBatchLipid) *gmodels.RecipeBatchLipid {
	if m == nil {
		return nil
	}

	r := &gmodels.RecipeBatchLipid{
		ID:     RecipeBatchLipidIDToGraphQL(uint(m.ID)),
		Weight: m.Weight,
		Cost:   m.Cost, Lipid: LipidIDToGraphQL(uint(m.LipidID)),
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
	}

	if boilergql.IntIsFilled(m.LipidID) {
		if m.R != nil && m.R.Lipid != nil {
			r.Lipid = LipidToGraphQL(m.R.Lipid)
		} else {
			r.Lipid = LipidWithIntID(m.LipidID)
		}
	}

	if boilergql.IntIsFilled(m.BatchID) {
		if m.R != nil && m.R.Batch != nil {
			r.Batch = RecipeBatchToGraphQL(m.R.Batch)
		} else {
			r.Batch = RecipeBatchWithIntID(m.BatchID)
		}
	}

	return r
}

func RecipeBatchLipidID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeBatchLipidIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeBatchLyeWithUintID(id uint) *gmodels.RecipeBatchLye {
	return &gmodels.RecipeBatchLye{
		ID: RecipeBatchLyeIDToGraphQL(id),
	}
}

func RecipeBatchLyeWithIntID(id int) *gmodels.RecipeBatchLye {
	return RecipeBatchLyeWithUintID(uint(id))
}

func RecipeBatchLyeWithNullDotUintID(id null.Uint) *gmodels.RecipeBatchLye {
	return RecipeBatchLyeWithUintID(id.Uint)
}

func RecipeBatchLyeWithNullDotIntID(id null.Int) *gmodels.RecipeBatchLye {
	return RecipeBatchLyeWithUintID(uint(id.Int))
}

func RecipeBatchLyesToGraphQL(am []*models.RecipeBatchLye) []*gmodels.RecipeBatchLye {
	ar := make([]*gmodels.RecipeBatchLye, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchLyeToGraphQL(m)
	}
	return ar
}

func RecipeBatchLyeIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeBatchLye)
}

func RecipeBatchLyeToGraphQL(m *models.RecipeBatchLye) *gmodels.RecipeBatchLye {
	if m == nil {
		return nil
	}

	r := &gmodels.RecipeBatchLye{
		ID:        RecipeBatchLyeIDToGraphQL(uint(m.ID)),
		Weight:    m.Weight,
		Discount:  m.Discount,
		Cost:      m.Cost,
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
	}

	if boilergql.IntIsFilled(m.LyeID) {
		if m.R != nil && m.R.Lye != nil {
			r.Lye = LyeToGraphQL(m.R.Lye)
		} else {
			r.Lye = LyeWithIntID(m.LyeID)
		}
	}

	if boilergql.IntIsFilled(m.BatchID) {
		if m.R != nil && m.R.Batch != nil {
			r.Batch = RecipeBatchToGraphQL(m.R.Batch)
		} else {
			r.Batch = RecipeBatchWithIntID(m.BatchID)
		}
	}

	return r
}

func RecipeBatchLyeID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeBatchLyeIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeBatchNoteWithUintID(id uint) *gmodels.RecipeBatchNote {
	return &gmodels.RecipeBatchNote{
		ID: RecipeBatchNoteIDToGraphQL(id),
	}
}

func RecipeBatchNoteWithIntID(id int) *gmodels.RecipeBatchNote {
	return RecipeBatchNoteWithUintID(uint(id))
}

func RecipeBatchNoteWithNullDotUintID(id null.Uint) *gmodels.RecipeBatchNote {
	return RecipeBatchNoteWithUintID(id.Uint)
}

func RecipeBatchNoteWithNullDotIntID(id null.Int) *gmodels.RecipeBatchNote {
	return RecipeBatchNoteWithUintID(uint(id.Int))
}

func RecipeBatchNotesToGraphQL(am []*models.RecipeBatchNote) []*gmodels.RecipeBatchNote {
	ar := make([]*gmodels.RecipeBatchNote, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchNoteToGraphQL(m)
	}
	return ar
}

func RecipeBatchNoteIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeBatchNote)
}

func RecipeBatchNoteToGraphQL(m *models.RecipeBatchNote) *gmodels.RecipeBatchNote {
	if m == nil {
		return nil
	}

	r := &gmodels.RecipeBatchNote{
		ID:        RecipeBatchNoteIDToGraphQL(uint(m.ID)),
		Note:      m.Note,
		Link:      m.Link,
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.BatchID) {
		if m.R != nil && m.R.Batch != nil {
			r.Batch = RecipeBatchToGraphQL(m.R.Batch)
		} else {
			r.Batch = RecipeBatchWithIntID(m.BatchID)
		}
	}

	return r
}

func RecipeBatchNoteID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeBatchNoteIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeFragranceWithUintID(id uint) *gmodels.RecipeFragrance {
	return &gmodels.RecipeFragrance{
		ID: RecipeFragranceIDToGraphQL(id),
	}
}

func RecipeFragranceWithIntID(id int) *gmodels.RecipeFragrance {
	return RecipeFragranceWithUintID(uint(id))
}

func RecipeFragranceWithNullDotUintID(id null.Uint) *gmodels.RecipeFragrance {
	return RecipeFragranceWithUintID(id.Uint)
}

func RecipeFragranceWithNullDotIntID(id null.Int) *gmodels.RecipeFragrance {
	return RecipeFragranceWithUintID(uint(id.Int))
}

func RecipeFragrancesToGraphQL(am []*models.RecipeFragrance) []*gmodels.RecipeFragrance {
	ar := make([]*gmodels.RecipeFragrance, len(am))
	for i, m := range am {
		ar[i] = RecipeFragranceToGraphQL(m)
	}
	return ar
}

func RecipeFragranceIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeFragrance)
}

func RecipeFragranceToGraphQL(m *models.RecipeFragrance) *gmodels.RecipeFragrance {
	if m == nil {
		return nil
	}

	r := &gmodels.RecipeFragrance{
		ID:         RecipeFragranceIDToGraphQL(uint(m.ID)),
		Percentage: m.Percentage,
		DeletedAt:  boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		CreatedAt:  boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:  boilergql.TimeDotTimeToInt(m.UpdatedAt),
	}

	if boilergql.IntIsFilled(m.FragranceID) {
		if m.R != nil && m.R.Fragrance != nil {
			r.Fragrance = FragranceToGraphQL(m.R.Fragrance)
		} else {
			r.Fragrance = FragranceWithIntID(m.FragranceID)
		}
	}

	if boilergql.IntIsFilled(m.RecipeID) {
		if m.R != nil && m.R.Recipe != nil {
			r.Recipe = RecipeToGraphQL(m.R.Recipe)
		} else {
			r.Recipe = RecipeWithIntID(m.RecipeID)
		}
	}

	return r
}

func RecipeFragranceID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeFragranceIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeLipidWithUintID(id uint) *gmodels.RecipeLipid {
	return &gmodels.RecipeLipid{
		ID: RecipeLipidIDToGraphQL(id),
	}
}

func RecipeLipidWithIntID(id int) *gmodels.RecipeLipid {
	return RecipeLipidWithUintID(uint(id))
}

func RecipeLipidWithNullDotUintID(id null.Uint) *gmodels.RecipeLipid {
	return RecipeLipidWithUintID(id.Uint)
}

func RecipeLipidWithNullDotIntID(id null.Int) *gmodels.RecipeLipid {
	return RecipeLipidWithUintID(uint(id.Int))
}

func RecipeLipidsToGraphQL(am []*models.RecipeLipid) []*gmodels.RecipeLipid {
	ar := make([]*gmodels.RecipeLipid, len(am))
	for i, m := range am {
		ar[i] = RecipeLipidToGraphQL(m)
	}
	return ar
}

func RecipeLipidIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeLipid)
}

func RecipeLipidToGraphQL(m *models.RecipeLipid) *gmodels.RecipeLipid {
	if m == nil {
		return nil
	}

	r := &gmodels.RecipeLipid{
		ID:         RecipeLipidIDToGraphQL(uint(m.ID)),
		Percentage: m.Percentage, Lipid: LipidIDToGraphQL(uint(m.LipidID)),
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
	}

	if boilergql.IntIsFilled(m.LipidID) {
		if m.R != nil && m.R.Lipid != nil {
			r.Lipid = LipidToGraphQL(m.R.Lipid)
		} else {
			r.Lipid = LipidWithIntID(m.LipidID)
		}
	}

	if boilergql.IntIsFilled(m.RecipeID) {
		if m.R != nil && m.R.Recipe != nil {
			r.Recipe = RecipeToGraphQL(m.R.Recipe)
		} else {
			r.Recipe = RecipeWithIntID(m.RecipeID)
		}
	}

	return r
}

func RecipeLipidID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeLipidIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeStepWithUintID(id uint) *gmodels.RecipeStep {
	return &gmodels.RecipeStep{
		ID: RecipeStepIDToGraphQL(id),
	}
}

func RecipeStepWithIntID(id int) *gmodels.RecipeStep {
	return RecipeStepWithUintID(uint(id))
}

func RecipeStepWithNullDotUintID(id null.Uint) *gmodels.RecipeStep {
	return RecipeStepWithUintID(id.Uint)
}

func RecipeStepWithNullDotIntID(id null.Int) *gmodels.RecipeStep {
	return RecipeStepWithUintID(uint(id.Int))
}

func RecipeStepsToGraphQL(am []*models.RecipeStep) []*gmodels.RecipeStep {
	ar := make([]*gmodels.RecipeStep, len(am))
	for i, m := range am {
		ar[i] = RecipeStepToGraphQL(m)
	}
	return ar
}

func RecipeStepIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeStep)
}

func RecipeStepToGraphQL(m *models.RecipeStep) *gmodels.RecipeStep {
	if m == nil {
		return nil
	}

	r := &gmodels.RecipeStep{
		ID:        RecipeStepIDToGraphQL(uint(m.ID)),
		Num:       m.Num,
		Note:      m.Note,
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.RecipeID) {
		if m.R != nil && m.R.Recipe != nil {
			r.Recipe = RecipeToGraphQL(m.R.Recipe)
		} else {
			r.Recipe = RecipeWithIntID(m.RecipeID)
		}
	}

	return r
}

func RecipeStepID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeStepIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func SupplierWithUintID(id uint) *gmodels.Supplier {
	return &gmodels.Supplier{
		ID: SupplierIDToGraphQL(id),
	}
}

func SupplierWithIntID(id int) *gmodels.Supplier {
	return SupplierWithUintID(uint(id))
}

func SupplierWithNullDotUintID(id null.Uint) *gmodels.Supplier {
	return SupplierWithUintID(id.Uint)
}

func SupplierWithNullDotIntID(id null.Int) *gmodels.Supplier {
	return SupplierWithUintID(uint(id.Int))
}

func SuppliersToGraphQL(am []*models.Supplier) []*gmodels.Supplier {
	ar := make([]*gmodels.Supplier, len(am))
	for i, m := range am {
		ar[i] = SupplierToGraphQL(m)
	}
	return ar
}

func SupplierIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.Supplier)
}

func SupplierToGraphQL(m *models.Supplier) *gmodels.Supplier {
	if m == nil {
		return nil
	}

	r := &gmodels.Supplier{
		ID:        SupplierIDToGraphQL(uint(m.ID)),
		Name:      m.Name,
		Website:   m.Website,
		Note:      m.Note,
		DeletedAt: boilergql.NullDotTimeToPointerInt(m.DeletedAt),
		UpdatedAt: boilergql.TimeDotTimeToInt(m.UpdatedAt),
		CreatedAt: boilergql.TimeDotTimeToInt(m.CreatedAt),
	}

	return r
}

func SupplierID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func SupplierIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

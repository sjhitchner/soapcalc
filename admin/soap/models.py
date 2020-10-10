import uuid
import string
import random
from django.db import models


class BaseModel(models.Model):

    class Meta:
        abstract = True

    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)
    deleted_at = models.DateTimeField(null=True)  


class InventoryModel(models.Model):

    class Meta:
        abstract = True

    supplier = models.ForeignKey('Supplier', on_delete=models.PROTECT)
    purchase_date = models.DateTimeField()
    expiry_date = models.DateTimeField()
    cost = models.FloatField()
    weight = models.FloatField()


class Supplier(BaseModel):

    class Meta:
        db_table = 'supplier'

    name = models.CharField(max_length=100)
    website = models.CharField(max_length=255)
    note = models.TextField(blank=True)


class Additive(BaseModel):

    class Meta:
        db_table = 'additive'

    name = models.CharField(max_length=100)
    note = models.TextField(blank=True)


class AdditiveInventory(BaseModel, InventoryModel):

    class Meta:
        db_table = 'additive_inventory'

    additive = models.ForeignKey('Additive', on_delete=models.PROTECT)


class Fragrance(BaseModel):

    class Meta:
        db_table = 'fragrance'

    name = models.CharField(max_length=100)
    note = models.TextField(blank=True)


class FragranceInventory(BaseModel, InventoryModel):

    class Meta:
        db_table = 'fragrance_inventory'

    fragrance = models.ForeignKey('Fragrance', on_delete=models.PROTECT)


class Lipid(BaseModel):

    class Meta:
        db_table = 'lipid'

    name = models.CharField(max_length=100)
    lauric = models.IntegerField()
    myristic = models.IntegerField()
    palmitic = models.IntegerField()
    stearic = models.IntegerField()
    ricinoleic = models.IntegerField()
    oleic = models.IntegerField()
    linoleic = models.IntegerField()
    linolenic = models.IntegerField()
    hardness = models.IntegerField()
    cleansing = models.IntegerField()
    conditioning = models.IntegerField()
    bubbly = models.IntegerField()
    creamy = models.IntegerField()
    iodine = models.IntegerField()
    ins = models.IntegerField()
    inci_name = models.CharField(max_length=100)


class LipidInventory(BaseModel, InventoryModel):

    class Meta:
        db_table = 'lipid_inventory'

    lipid = models.ForeignKey('Lipid', on_delete=models.PROTECT)    
    sap = models.FloatField() 
    naoh = models.FloatField()
    koh = models.FloatField()
    grams_per_liter = models.FloatField()


class Lye(BaseModel):

    class Meta:
        db_table = 'lye'

    kind = models.CharField(max_length=4, choices=(('naoh', 'NaOH'), ('koh', 'KOH')), default='naoh')
    name = models.CharField(max_length=100)
    note = models.TextField(blank=True)


class LyeInventory(BaseModel, InventoryModel):

    class Meta:
        db_table = 'lye_inventory'

    lye = models.ForeignKey('LyeInventory', on_delete=models.PROTECT)
    concentration = models.FloatField()


class Recipe(BaseModel):

    class Meta:
        db_table = 'recipe'

    name = models.CharField(max_length=100)
    note = models.TextField(blank=True)


class RecipeAdditive(BaseModel):

    class Meta:
        db_table = 'recipe_additive'

    recipe = models.ForeignKey('Recipe', on_delete=models.PROTECT)
    additive = models.ForeignKey('Additive', on_delete=models.PROTECT)
    percentage = models.FloatField()


class RecipeFragrance(BaseModel):

    class Meta:
        db_table = 'recipe_fragrance'

    recipe = models.ForeignKey('Recipe', on_delete=models.PROTECT)
    fragrance = models.ForeignKey('Fragrance', on_delete=models.PROTECT)
    percentage = models.FloatField()


class RecipeLipid(BaseModel):

    class Meta:
        db_table = 'recipe_lipid'

    recipe = models.ForeignKey('Recipe', on_delete=models.PROTECT)
    lipid = models.ForeignKey('Lipid', on_delete=models.PROTECT)
    percentage = models.FloatField()


class RecipeBatch(BaseModel):

    class Meta:
        db_table = 'recipe_batch'

    recipe = models.ForeignKey('Recipe', on_delete=models.PROTECT)
    batch_id = models.CharField(max_length=16)
    production_date = models.DateTimeField()
    cured_produced = models.DateTimeField()
    note = models.TextField(blank=True)

    lipid_weight = models.FloatField()
    production_weight = models.FloatField()
    cured_weight = models.FloatField()


class RecipeBatchLye(BaseModel):

    class Meta:
        db_table = 'recipe_batch_lye'

    lye = models.ForeignKey('Lye', on_delete=models.PROTECT)
    weight = models.FloatField()
    discount = models.FloatField()
    cost = models.FloatField()


class RecipeBatchAdditive(BaseModel):

    class Meta:
        db_table = 'recipe_batch_additive'

    additive = models.ForeignKey('Additive', on_delete=models.PROTECT)
    weight = models.FloatField()
    cost = models.FloatField()


class RecipeBatchFragrance(BaseModel):

    class Meta:
        db_table = 'recipe_batch_fragrance'

    fragrance = models.ForeignKey('Fragrance', on_delete=models.PROTECT)
    weight = models.FloatField()
    cost = models.FloatField()


class RecipeBatchLipid(BaseModel):

    class Meta:
        db_table = 'recipe_batch_lipid'

    lipid = models.ForeignKey('Lipid', on_delete=models.PROTECT)
    weight = models.FloatField()
    cost = models.FloatField()

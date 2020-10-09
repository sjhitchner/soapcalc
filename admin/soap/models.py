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


class Additive(BaseModel):
    name = models.CharField(max_length=100)
#    note = models.TextField(blank=True)


class AdditiveInventory(BaseModel, InventoryModel):
    additive = models.ForeignKey('Additive', on_delete=models.PROTECT)


class Fragrance(BaseModel):
    name = models.CharField(max_length=100)
    grams_per_liter = models.FloatField()
    note = models.TextField(blank=True)


class FragranceInventory(BaseModel, InventoryModel):
    fragrance = models.ForeignKey('Fragrance', on_delete=models.PROTECT)


class Supplier(BaseModel):
    name = models.CharField(max_length=100)
    website = models.CharField(max_length=255)
    note = models.TextField(blank=True)


class Lipid(BaseModel):
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
    lipid = models.ForeignKey('Lipid', on_delete=models.PROTECT)    
    sap = models.FloatField() 
    naoh = models.FloatField()
    koh = models.FloatField()
    grams_per_liter = models.FloatField()


class Lye(BaseModel):
    kind = models.CharField(max_length=4, choices=(('naoh', 'NaOH'), ('koh', 'KOH')), default='naoh')
    name = models.CharField(max_length=100)
    note = models.TextField(blank=True)


class LyeInventory(BaseModel, InventoryModel):
    lye = models.ForeignKey('LyeInventory', on_delete=models.PROTECT)


class Recipe(BaseModel):
    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
    name = models.CharField(max_length=100)
    lye = models.ForeignKey('RecipeLye', on_delete=models.CASCADE)
    additives = models.ForeignKey('RecipeAdditive', on_delete=models.CASCADE)
    lipids = models.ForeignKey('RecipeLipid', on_delete=models.CASCADE)
    lipidweight = models.FloatField()
    fragrances = models.ForeignKey('RecipeFragrance', on_delete=models.CASCADE)
    cost = models.FloatField()
    note = models.TextField(blank=True)


class RecipeAdditive(BaseModel):
    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
    additive = models.ForeignKey('Additive', on_delete=models.PROTECT)
    percentage = models.FloatField()
    weight = models.FloatField()
    cost = models.FloatField()


class RecipeFragrance(BaseModel):
    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
    fragrance = models.ForeignKey('Fragrance', on_delete=models.PROTECT)
    percentage = models.FloatField()
    weight = models.FloatField()


class RecipeLipid(BaseModel):
    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
    lipid = models.ForeignKey('Lipid', on_delete=models.PROTECT)
    percentage = models.FloatField()
    weight = models.FloatField()
    cost = models.FloatField()


class RecipeLye(BaseModel):
    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
    lye = models.ForeignKey('Lye', on_delete=models.PROTECT)
    weight = models.FloatField()
    concentration = models.FloatField()
    discount = models.FloatField()
    cost = models.FloatField()

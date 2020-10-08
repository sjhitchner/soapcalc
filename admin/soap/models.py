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


class Additive(BaseModel):
    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
    name = models.CharField(max_length=100)
    inventory = models.ForeignKey('Inventory', on_delete=models.PROTECT)
    note = models.TextField(blank=True)


class Fragrance(BaseModel):
    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
    name = models.CharField(max_length=100)
    grams_per_liter = models.FloatField()
    inventory = models.ForeignKey('Inventory', on_delete=models.PROTECT)
    note = models.TextField(blank=True)


class Inventory(BaseModel):
    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
    supplier = models.ForeignKey('Supplier', on_delete=models.PROTECT)
    purchase_date = models.DateTimeField()
    expiry_date = models.DateTimeField()
    cost = models.FloatField()
    weight = models.FloatField()

# class InventoryCategory(BaseModel):
#    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
#    category = models.CharField(max_length=30)


class Supplier(BaseModel):
    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
    name = models.CharField(max_length=100)
    website = models.CharField(max_length=255)
    note = models.TextField(blank=True)


class Lipid(BaseModel):

    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
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


class LipidBatch(BaseModel):
    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
    lipid = models.ForeignKey('Lipid', on_delete=models.PROTECT)    
    naoh = models.FloatField()
    koh = models.FloatField()
    grams_per_liter = models.FloatField()
    inventory = models.ForeignKey('Inventory', on_delete=models.PROTECT)


class Lye(BaseModel):

    id = models.CharField(primary_key=True, max_length=36, default=uuid.uuid1)
    kind = models.CharField(max_length=4, choices=(('naoh', 'NaOH'), ('koh', 'KOH')), default='naoh')
    name = models.CharField(max_length=100)
    note = models.TextField(blank=True)
    inventory = models.ForeignKey('Inventory', on_delete=models.PROTECT)


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
    lipid = models.ForeignKey('LipidBatch', on_delete=models.PROTECT)
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


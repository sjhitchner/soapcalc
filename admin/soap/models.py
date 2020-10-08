from django.db import models


class Additive(models.Model):

    id = models.CharField(max_length=100)
    name = models.CharField(max_length=100)
    note = models.CharField(max_length=100)
    inventory = models.ForeignKey()



class Fragrance(models.Model):

    id = models.CharField(max_length=100)
    name = models.CharField(max_length=100)
    gramsperliter = models.FloatField()
    inventory = models.ForeignKey()
    note = models.CharField(max_length=100)



class Inventory(models.Model):

    id = models.CharField(max_length=100)
    supplier = models.ForeignKey()
    purchasedate = models.ForeignKey()
    expirydate = models.ForeignKey()
    cost = models.FloatField()
    weight = models.FloatField()



class Lipid(models.Model):

    id = models.CharField(max_length=100)
    name = models.CharField(max_length=100)
    type = models.CharField(max_length=100)
    naoh = models.FloatField()
    koh = models.FloatField()
    inciname = models.CharField(max_length=100)
    gramsperliter = models.FloatField()
    inventory = models.ForeignKey()



class Lye(models.Model):

    id = models.CharField(max_length=100)
    name = models.CharField(max_length=100)
    note = models.CharField(max_length=100)
    inventory = models.ForeignKey()



class Recipe(models.Model):

    id = models.CharField(max_length=100)
    name = models.CharField(max_length=100)
    lye = models.ForeignKey()
    additives = models.ForeignKey()
    lipids = models.ForeignKey()
    lipidweight = models.FloatField()
    fragrances = models.ForeignKey()
    quality = models.ForeignKey()
    composition = models.ForeignKey()
    cost = models.FloatField()



class RecipeAdditive(models.Model):

    id = models.CharField(max_length=100)
    additive = models.ForeignKey()
    percentage = models.FloatField()
    weight = models.FloatField()
    cost = models.FloatField()



class RecipeComposition(models.Model):

    id = models.CharField(max_length=100)
    lauric = models.IntegerField()
    myristic = models.IntegerField()
    palmitic = models.IntegerField()
    stearic = models.IntegerField()
    ricinoleic = models.IntegerField()
    oleic = models.IntegerField()
    linoleic = models.IntegerField()
    linolenic = models.IntegerField()



class RecipeFragrance(models.Model):

    id = models.CharField(max_length=100)
    fragrance = models.ForeignKey()
    percentage = models.FloatField()
    weight = models.FloatField()



class RecipeLipid(models.Model):

    id = models.CharField(max_length=100)
    lipid = models.ForeignKey()
    percentage = models.FloatField()
    weight = models.FloatField()
    cost = models.FloatField()



class RecipeLye(models.Model):

    id = models.CharField(max_length=100)
    type = models.ForeignKey()
    weight = models.FloatField()
    concentration = models.FloatField()
    discount = models.FloatField()
    cost = models.FloatField()



class RecipeQuality(models.Model):

    id = models.CharField(max_length=100)
    hardness = models.IntegerField()
    cleansing = models.IntegerField()
    conditioning = models.IntegerField()
    bubbly = models.IntegerField()
    creamy = models.IntegerField()
    iodine = models.IntegerField()
    ins = models.IntegerField()



class Supplier(models.Model):

    id = models.CharField(max_length=100)
    name = models.CharField(max_length=100)
    note = models.CharField(max_length=100)



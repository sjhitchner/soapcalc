from django.contrib import admin
from django.contrib.admin import DateFieldListFilter
from .models import (
    Additive,
    Fragrance,
    Inventory,
    Supplier,
    Lipid,
    LipidBatch,
    Lye,
    Recipe,
    RecipeAdditive,
    RecipeFragrance,
    RecipeLipid,
    RecipeLye,
)


class AdditiveAdmin(admin.ModelAdmin):
    pass

class FragranceAdmin(admin.ModelAdmin):
    pass

class InventoryAdmin(admin.ModelAdmin):
    pass

class SupplierAdmin(admin.ModelAdmin):
    pass

class LipidAdmin(admin.ModelAdmin):
    pass

class LipidBatchAdmin(admin.ModelAdmin):
    pass

class LyeAdmin(admin.ModelAdmin):
    pass

class RecipeAdmin(admin.ModelAdmin):
    pass

class RecipeAdditiveAdmin(admin.ModelAdmin):
    pass

class RecipeFragranceAdmin(admin.ModelAdmin):
    pass

class RecipeLipidAdmin(admin.ModelAdmin):
    pass

class RecipeLyeAdmin(admin.ModelAdmin):
    pass


admin.site.register(Additive, AdditiveAdmin)
admin.site.register(Fragrance, FragranceAdmin)
admin.site.register(Inventory, InventoryAdmin)
admin.site.register(Supplier, SupplierAdmin)
admin.site.register(Lipid, LipidAdmin)
admin.site.register(LipidBatch, LipidBatchAdmin)
admin.site.register(Lye, LyeAdmin)
admin.site.register(Recipe, RecipeAdmin)
admin.site.register(RecipeAdditive, RecipeAdditiveAdmin)
admin.site.register( RecipeFragrance, RecipeFragranceAdmin)
admin.site.register(RecipeLipid, RecipeLipidAdmin)
admin.site.register(RecipeLye, RecipeLyeAdmin)

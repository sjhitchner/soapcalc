# Generated by Django 3.1.2 on 2020-10-10 23:04

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ('soap', '0002_auto_20201010_2118'),
    ]

    operations = [
        migrations.RenameField(
            model_name='recipebatch',
            old_name='batch_id',
            new_name='batch',
        ),
    ]